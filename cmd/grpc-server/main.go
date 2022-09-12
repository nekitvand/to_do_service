package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nekitvand/to_do_service/internal/server"
	"github.com/nekitvand/to_do_service/internal/config"
	todo_service "github.com/nekitvand/to_do_service/internal/service"
	"github.com/nekitvand/to_do_service/internal/pkg/db"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

)

func main() {

	config.ReadConfigYaml("config.yaml")
	cfg := config.GetConfig()
	

	todoRepository := todo_service.NewRepository()
	todoService := todo_service.NewService(todoRepository)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func ()  {
		server.GoGatewayRpcRun(&cfg)
	}()

	go func() {
		gserv := server.NewGrpcServer(todoService,)
		gserv.GoRpcRun(&cfg)
	}()

	go func() {
		swagger,err := server.CreateSwaggerServer(&cfg)
		if err != nil{
			fmt.Println(err)
		}
		swagger.ListenAndServe()
	}()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		fmt.Println(fmt.Errorf("sql.Open() error | %v,",err))
	}
	defer conn.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		fmt.Println("Выход c",v)
	case done := <-ctx.Done():
		fmt.Println("ctx.Done:", done)
	}

}