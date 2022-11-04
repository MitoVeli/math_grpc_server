package main

import (
	"fmt"
	"net/http"

	grpc "github.com/MitoVeli/math_grpc_server/grpc"
	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	mathOperationsService := mathOperationsService.NewMathOperationsService()

	// TODO: to be deleted
	var result int32
	mathOperationsService.Add(1, 2, &result)

	fmt.Println("the result is", result)
	////////////

	/* ------------- INITIALIZE gRPC ------------- */
	go grpc.GrpcServer()

	s := http.Server{
		Addr: ":8080",
	}

	s.ListenAndServe()

	// TODO: add variables to config file
	// TODO: built client app for grpc
	// TODO: add tests
	// TODO: add rest of the operations

}
