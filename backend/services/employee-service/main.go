package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	pb "github.com/tkp-VL037/go-employee/proto"
	"github.com/tkp-VL037/go-employee/services/employee-service/api"
	"github.com/tkp-VL037/go-employee/services/employee-service/db"
	"github.com/tkp-VL037/go-employee/services/employee-service/nsq"
)

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
	pb.RegisterEmployeeServiceServer(srv, api.EmployeeServer{})

	log.Println("Starting RPC server at", serviceEmployeePort)

	l, err := net.Listen("tcp", serviceEmployeePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", serviceEmployeePort, err)
	}

	log.Fatal(srv.Serve(l))
}
