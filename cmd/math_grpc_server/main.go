package main

import (
	"log"
	"net/http"

	configs "github.com/MitoVeli/math_grpc_server/configs"

	grpc "github.com/MitoVeli/math_grpc_server/pkg/grpc"

	mathService "github.com/MitoVeli/math_grpc_server/pkg/math"
)

func main() {

	// initialize mathOperationsService
	mathOperationService := mathService.NewMathOperationsService()

	// initialize gRPC server
	mathServer := grpc.NewMathOperationsServiceServer(mathOperationService)
	go grpc.GrpcServer(mathServer)

	log.Printf("gRPC server started on port: %s", configs.GrpcPort)

	// start http server
	s := http.Server{
		Addr: configs.AppPort,
	}
	s.ListenAndServe()

}
