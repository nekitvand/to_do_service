package server


import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	service "github/nekitvand/to_do_service/pkg/to_do_service"

)


const grpcHostPort = "0.0.0.0:8082"


type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
	}


func (s *GrpcServer) GoRpcRun() error{
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		return err
	}
	defer func() {
		err := listen.Close()
		if err != nil{
			fmt.Println("Failed to close")
	}}()
		
	service.RegisterToDoServiceServer(grpcServer,nil)
	return grpcServer.Serve(listen)
}