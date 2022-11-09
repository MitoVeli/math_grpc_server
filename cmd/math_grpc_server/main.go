package main

import (
	"net/http"

	configs "github.com/MitoVeli/math_grpc_server/configs"

	grpc "github.com/MitoVeli/math_grpc_server/pkg/grpc"

	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg/math"
)

func main() {

	// initialize mathOperationsService
	mathOperationsService.NewMathOperationsService()

	// initialize gRPC server
	go grpc.GrpcServer()

	// start http server
	s := http.Server{
		Addr: configs.AppPort,
	}
	s.ListenAndServe()

}
