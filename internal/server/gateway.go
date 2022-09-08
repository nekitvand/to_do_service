package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	service "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	gatewayGrpcHostPort = "localhost:8080"
)

func GoGatewayRpcRun() error {

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := service.RegisterToDoServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		panic(err)
	}
	gwServer := &http.Server{
		Addr:    gatewayGrpcHostPort,
		Handler: cors(mux),
	}

	return gwServer.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providedOrigin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", providedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		h.ServeHTTP(w, r)
	})
}
