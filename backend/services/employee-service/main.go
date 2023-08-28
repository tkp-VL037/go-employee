package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/ksuid"
	"github.com/tkp-VL037/go-employee/model"
	pb "github.com/tkp-VL037/go-employee/proto"
	"github.com/tkp-VL037/go-employee/services/employee-service/constant"
	"github.com/tkp-VL037/go-employee/services/employee-service/db"
	"github.com/tkp-VL037/go-employee/services/employee-service/nsq"
)

type EmployeeServer struct {
	pb.UnimplementedEmployeeServiceServer
}

type Message struct {
	Sender    string
	Content   interface{}
	Timestamp time.Time
}

func (EmployeeServer) GetEmployees(ctx context.Context, param *pb.NoParam) (*pb.GetEmployeesResponse, error) {
	var employees []*model.Employee

	result := db.PostgresDB.Preload("Statistic").Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	employeesStatsResponse := make([]*pb.EmployeeResponse, len(employees))
	for i, es := range employees {
		employeesStatsResponse[i] = &pb.EmployeeResponse{
			Employee: &pb.Employee{
				Id:       es.ID,
				Name:     es.Name,
				Age:      int32(es.Age),
				Position: es.Position,
			},
			Statistic: &pb.Statistic{
				Id:         es.Statistic.ID,
				EmployeeID: es.Statistic.EmployeeID,
				ViewCount:  es.Statistic.ViewCount,
			},
		}
	}

	return &pb.GetEmployeesResponse{
		EmployeeResponse: employeesStatsResponse,
	}, nil
}

func (EmployeeServer) GetEmployeeDetail(ctx context.Context, param *pb.GetEmployeeDetailRequest) (*pb.EmployeeResponse, error) {
	var employee *model.Employee
	var employeeRes *pb.EmployeeResponse

	key := fmt.Sprintf(constant.FIND_ONE_EMPLOYEE, param.Id)
	db.RedisClient.Del(ctx, key) // delete cache KEY
	cachedEmployeeJSON, err := db.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err := db.PostgresDB.Preload("Statistic").First(&employee, "id = ?", param.Id).Error; err != nil {
			return nil, err
		}

		err = db.PostgresDB.Model(&model.Statistic{}).Where("employee_id = ?", param.Id).
			UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
		if err != nil {
			return nil, err
		}

		cache, err := json.Marshal(employee)
		fmt.Println("saving employee detail to cache...")
		err = db.RedisClient.Set(ctx, key, cache, redis.KeepTTL).Err()
		if err != nil {
			fmt.Println("FAILED TO SET CACHE", err)
			return nil, err
		}

		employeeRes = &pb.EmployeeResponse{
			Employee: &pb.Employee{
				Id:       employee.ID,
				Name:     employee.Name,
				Age:      int32(employee.Age),
				Position: employee.Position,
			}, Statistic: &pb.Statistic{
				Id:        employee.Statistic.ID,
				ViewCount: employee.Statistic.ViewCount,
			},
		}
	} else {
		if err := json.Unmarshal([]byte(cachedEmployeeJSON), &employee); err != nil {
			return nil, err
		}
		employeeRes = &pb.EmployeeResponse{
			Employee: &pb.Employee{
				Id:       employee.ID,
				Name:     employee.Name,
				Age:      int32(employee.Age),
				Position: employee.Position,
			}, Statistic: &pb.Statistic{
				Id:        employee.Statistic.ID,
				ViewCount: employee.Statistic.ViewCount,
			},
		}
		fmt.Println("employee detail retrieved from cache!")
	}

	// TODO: publish to NSQ
	topic := constant.TOPIC_EMPLOYEE_DETAIL
	msg := Message{
		Sender:    "EMPLOYEE_SERVICE",
		Content:   employeeRes,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = nsq.NsqProducer.Publish(topic, payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GetEmployeeDetail Message published!")

	return employeeRes, nil
}

func (EmployeeServer) AddEmployee(ctx context.Context, param *pb.AddEmployeeRequest) (*pb.EmployeeResponse, error) {
	employee := &model.Employee{
		ID:       ksuid.New().String(),
		Name:     param.Name,
		Age:      int(param.Age),
		Position: param.Position,
	}

	statistic := &model.Statistic{
		ID:         ksuid.New().String(),
		EmployeeID: employee.ID,
		ViewCount:  0,
	}

	err := db.PostgresDB.Create(employee).Error
	if err != nil {
		return nil, err
	}

	err = db.PostgresDB.Create(statistic).Error
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Employee: &pb.Employee{
			Id:       employee.ID,
			Name:     employee.Name,
			Age:      int32(employee.Age),
			Position: employee.Position,
		},
		Statistic: &pb.Statistic{
			Id:         statistic.ID,
			EmployeeID: statistic.EmployeeID,
			ViewCount:  statistic.ViewCount,
		},
	}, nil
}

func (EmployeeServer) UpdateEmployee(ctx context.Context, param *pb.UpdateEmployeeRequest) (*pb.EmployeeResponse, error) {
	var employee *model.Employee
	if err := db.PostgresDB.Preload("Statistic").First(&employee, "id = ?", param.Id).Error; err != nil {
		return nil, err
	}

	employee.Name = param.Name
	employee.Position = param.Position

	if err := db.PostgresDB.Save(&employee).Error; err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Employee: &pb.Employee{
			Id:       employee.ID,
			Name:     employee.Name,
			Age:      int32(employee.Age),
			Position: employee.Position,
		},
		Statistic: &pb.Statistic{
			Id:         employee.Statistic.ID,
			EmployeeID: employee.Statistic.EmployeeID,
			ViewCount:  employee.Statistic.ViewCount,
		},
	}, nil
}

func (EmployeeServer) DeleteEmployee(ctx context.Context, param *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	err := db.PostgresDB.Where("employee_id = ?", param.Id).Delete(&model.Statistic{}).Error
	if err != nil {
		return &pb.DeleteEmployeeResponse{
			Success: false,
		}, nil
	}

	err = db.PostgresDB.Where("id = ?", param.Id).Delete(&model.Employee{}).Error

	return &pb.DeleteEmployeeResponse{
		Success: err == nil,
	}, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serviceEmployeePort := os.Getenv("EMPLOYEE_GRPC_PORT")

	if err := db.ConnectPostgres(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.ConnectRedis()
	nsq.StartProducer()

	srv := grpc.NewServer()
	var employeeSrv EmployeeServer
	pb.RegisterEmployeeServiceServer(srv, employeeSrv)

	log.Println("Starting RPC server at", serviceEmployeePort)

	l, err := net.Listen("tcp", serviceEmployeePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", serviceEmployeePort, err)
	}

	log.Fatal(srv.Serve(l))
}
