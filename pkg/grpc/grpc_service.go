package grpc

import (
	context "context"
	"log"

	pb "github.com/MitoVeli/math_grpc_client/pkg/grpc/math"
	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg/math"
)

type MathOperationsServiceServer struct {
	pb.UnimplementedMathOperationsServer
	mathOperationsService mathOperationsService.MathOperations
}

func NewMathOperationsServiceServer(mathOperationsService mathOperationsService.MathOperations) *MathOperationsServiceServer {
	return &MathOperationsServiceServer{mathOperationsService: mathOperationsService}
}

func (s *MathOperationsServiceServer) Calculate(ctx context.Context, in *pb.OperationRequest) (*pb.OperationResponse, error) {
	var result float32

	if err := s.mathOperationsService.Calculate(in.X, in.Y, in.OperationSign, &result); err != nil {
		log.Printf("error occured while math operation: %f %s %f, error: %v", in.X, in.OperationSign, in.Y, err)
		return nil, err
	}

	// return user
	return &pb.OperationResponse{
		Result: result,
	}, nil
}
