package grpc

import (
	"log"
	"net"

	pb "github.com/MitoVeli/math_grpc_client/pkg/grpc/math"
	configs "github.com/MitoVeli/math_grpc_server/configs"
	"google.golang.org/grpc"
)

type Server struct {
	MathOperationsServiceServer
}

func NewGrpcServer() *Server {
	return &Server{}
}

func GrpcServer() {
	lis, err := net.Listen("tcp", configs.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterMathOperationsServer(s, NewGrpcServer())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}
