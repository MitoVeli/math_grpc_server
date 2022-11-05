package grpc

import (
	context "context"
	"errors"
	"log"

	mathOperationsService "math_grpc_server/pkg"

	pb "github.com/MitoVeli/math_grpc_client/grpc"
)

type MathOperationsServiceServer struct {
	pb.UnimplementedMathOperationsServer
	mathOperationsService mathOperationsService.MathOperations
}

func (s *MathOperationsServiceServer) Add(ctx context.Context, in *pb.OperationRequest) (*pb.OperationResponse, error) {

	if err := s.mathOperationsService.Add(in.A, in.B, &in.Result); err != nil {
		log.Println("error in Add method", err)
		return nil, errors.New("failed to add")
	}

	// return user
	return &pb.OperationResponse{
		Result: in.Result,
	}, nil
}
