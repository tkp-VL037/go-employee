package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeDetail not implemented")
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
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}

func (EmployeeServer) DeleteEmployee(ctx context.Context, param *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
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
