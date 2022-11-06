package grpc

import (
	context "context"
	"errors"
	"log"

	pb "github.com/MitoVeli/math_grpc_client/grpc"
	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg"
)

type MathOperationsServiceServer struct {
	pb.UnimplementedMathOperationsServer
	mathOperationsService mathOperationsService.MathOperations
}

func (s *MathOperationsServiceServer) Add(ctx context.Context, in *pb.OperationRequest) (*pb.OperationResponse, error) {

	if err := s.mathOperationsService.DoMath(in.X, in.Y, in.OperationSign, &in.Result); err != nil {
		log.Printf("error occured while math operation: %d %s %d, error: %v", in.X, in.OperationSign, in.Y, err)
		return nil, errors.New("error occured while math operation")
	}

	// return user
	return &pb.OperationResponse{
		Result: in.Result,
	}, nil
}
