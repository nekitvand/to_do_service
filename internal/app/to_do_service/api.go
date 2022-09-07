package todo_service

import (
	// "context"
	to_do "github.com/nekitvand/to_do_service/internal/service"
	desc "github.com/nekitvand/to_do_service/pkg/to_do_service"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

type Implementation struct {
	desc.UnimplementedToDoServiceServer
	todoService to_do.Service
}

func NewUsersService(todoService to_do.Service) desc.ToDoServiceServer {
	return &Implementation{todoService: todoService,}
}
