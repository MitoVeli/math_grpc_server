package main

import (
	"log"
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

	log.Printf("gRPC server started on port: %s", configs.GrpcPort)

	// start http server
	s := http.Server{
		Addr: configs.AppPort,
	}
	s.ListenAndServe()

}
