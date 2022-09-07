package todo_service

import (
	"context"
	proto_to_do_service "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetToDoById(ctx context.Context, req *proto_to_do_service.GetToDoByIdRequest) (*proto_to_do_service.GetToDoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToDo not implemented")
}