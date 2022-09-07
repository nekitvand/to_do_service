package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	service "github/nekitvand/to_do_service/pkg/to_do_service"
)

const (
	gatewayGrpcHostPort = "0.0.0.0:8080"
)

func GoGatewayRpcRun() error {

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := service.RegisterToDoServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		panic(err)
	}

	return http.ListenAndServe(gatewayGrpcHostPort, mux)
}