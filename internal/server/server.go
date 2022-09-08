package server

import (
	"fmt"
	"net"

	// "net/http"

	api "github.com/nekitvand/to_do_service/internal/app/to_do_service"
	todo_service "github.com/nekitvand/to_do_service/internal/service"
	service "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"google.golang.org/grpc"
)

const grpcHostPort = "localhost:8082"

type GrpcServer struct {
	todoService todo_service.Service
}

func NewGrpcServer(todoService todo_service.Service) *GrpcServer {
	return &GrpcServer{todoService: todoService}
}

func (s *GrpcServer) GoRpcRun() error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		return err
	}
	defer func() {
		err := listen.Close()
		if err != nil {
			fmt.Println("Failed to close")
		}
	}()

	service.RegisterToDoServiceServer(grpcServer, api.NewUsersService(s.todoService))
	return grpcServer.Serve(listen)
}
