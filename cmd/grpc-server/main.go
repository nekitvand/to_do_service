package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nekitvand/to_do_service/internal/server"
	todo_service "github.com/nekitvand/to_do_service/internal/service"
)

func main() {

	todoRepository := todo_service.NewRepository()
	todoService := todo_service.NewService(todoRepository)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func ()  {
		server.GoGatewayRpcRun()
	}()

	go func() {
		gserv := server.NewGrpcServer(todoService,)
		gserv.GoRpcRun()
	}()

	go func() {
		swagger,err := server.CreateSwaggerServer("localhost:8080","localhost:6001","../../swagger/api/to_do_service/to_do_service.swagger.json")
		if err != nil{
			fmt.Println(err)
		}
		swagger.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		fmt.Println("Выход c",v)
	case done := <-ctx.Done():
		fmt.Println("ctx.Done:", done)
	}

}