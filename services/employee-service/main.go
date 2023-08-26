package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/segmentio/ksuid"
	"github.com/tkp-VL037/go-employee/db"
	"github.com/tkp-VL037/go-employee/model"
	pb "github.com/tkp-VL037/go-employee/proto"
)

type EmployeeServer struct {
	pb.UnimplementedEmployeeServiceServer
}

func (EmployeeServer) GetEmployees(ctx context.Context, param *pb.NoParam) (*pb.GetEmployeesResponse, error) {
	var employees []*model.Employee

	result := db.DB.Preload("Statistic").Find(&employees)
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
	if err := db.DB.Preload("Statistic").First(&employee, "id = ?", param.Id).Error; err != nil {
		return nil, err
	}

	err := db.DB.Model(&model.Statistic{}).Where("employee_id = ?", param.Id).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
	if err != nil {
		return nil, err
	}

	// TODO: publish to NSQ

	return &pb.EmployeeResponse{
		Employee: &pb.Employee{
			Id:       employee.ID,
			Name:     employee.Name,
			Age:      int32(employee.Age),
			Position: employee.Position,
		}, Statistic: &pb.Statistic{
			Id:        employee.Statistic.ID,
			ViewCount: employee.Statistic.ViewCount,
		},
	}, nil
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

	err := db.DB.Create(employee).Error
	if err != nil {
		return nil, err
	}

	err = db.DB.Create(statistic).Error
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
	if err := db.DB.Preload("Statistic").First(&employee, "id = ?", param.Id).Error; err != nil {
		return nil, err
	}

	employee.Name = param.Name
	employee.Position = param.Position

	if err := db.DB.Save(&employee).Error; err != nil {
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
	err := db.DB.Where("employee_id = ?", param.Id).Delete(&model.Statistic{}).Error
	if err != nil {
		return &pb.DeleteEmployeeResponse{
			Success: false,
		}, nil
	}

	err = db.DB.Where("id = ?", param.Id).Delete(&model.Employee{}).Error

	return &pb.DeleteEmployeeResponse{
		Success: err == nil,
	}, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serviceEmployeePort := os.Getenv("SERVICE_EMPLOYEE_PORT")

	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

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
