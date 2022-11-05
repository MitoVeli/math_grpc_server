package main

import (
	"net/http"

	configs "math_grpc_server/configs"

	grpc "math_grpc_server/grpc"

	mathOperationsService "math_grpc_server/pkg"
)

func main() {

	// TODO: add variables to config file
	// TODO: built client app for grpc
	// TODO: add tests
	// TODO: add rest of the operations
	// TODO: modify project structre as asked

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
