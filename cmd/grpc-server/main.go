package main

import (
	"context"
	"fmt"
	"github/nekitvand/to_do_service/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	go func ()  {
		server.GoGatewayRpcRun()
	}()

	go func() {
		gserv := server.NewGrpcServer()
		gserv.GoRpcRun()
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