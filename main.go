package main

import (
	"net/http"

	grpc "github.com/MitoVeli/math_grpc_server/grpc"
	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg"
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
		Addr: ":8080",
	}
	s.ListenAndServe()

}
