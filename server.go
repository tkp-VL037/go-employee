package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tkp-VL037/go-employee/constant"
	"github.com/tkp-VL037/go-employee/graph"
	"github.com/tkp-VL037/go-employee/graph/nsq"
	pb "github.com/tkp-VL037/go-employee/proto"
)

const defaultPort = "4000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			EmployeeSrvClient: employeeService(),
		},
	}))

	startConsumers()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func employeeService() pb.EmployeeServiceClient {
	port := os.Getenv("SERVICE_EMPLOYEE_PORT")
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to ", port, err)
	}

	return pb.NewEmployeeServiceClient(conn)
}

func startConsumers() {
	nsq.StartConsumer(constant.TOPIC_EMPLOYEE_DETAIL, constant.CHANNEL_PROCESSING)
	//
	//
}
