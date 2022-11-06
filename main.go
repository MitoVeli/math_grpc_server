package main

import (
	"net/http"

	configs "github.com/MitoVeli/math_grpc_server/configs"

	grpc "github.com/MitoVeli/math_grpc_server/grpc"

	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// TODO: tests
	// TODO: check config
	// TODO: check project structure, server, docker etc
	// TODO: check folder structure as asked!
	// TODO: check all comments and add if necessary

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
