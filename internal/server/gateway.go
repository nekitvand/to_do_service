package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cfg "github.com/nekitvand/to_do_service/internal/config"
	service "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GoGatewayRpcRun(cfg *cfg.Config) error {

	GrcAddr := fmt.Sprintf("%s:%v",cfg.Grpc.Host,cfg.Grpc.Port)
	GatewayAddr := fmt.Sprintf("%s:%v",cfg.Gateway.Host,cfg.Gateway.Port)

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := service.RegisterToDoServiceHandlerFromEndpoint(context.Background(), mux, GrcAddr, opts)
	if err != nil {
		panic(err)
	}
	gwServer := &http.Server{
		Addr:    GatewayAddr,
		Handler: cors(mux),
	}

	return gwServer.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT, PATCH, DELETE,OPTIONS",)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS"{
			w.WriteHeader(http.StatusOK)
		}
		h.ServeHTTP(w, r)
	})
}
