package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	pb "github.com/tkp-VL037/go-employee/proto"
)

type Resolver struct {
	EmployeeSrvClient pb.EmployeeServiceClient
}
