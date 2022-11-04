package grpc

import (
	context "context"
	"errors"

	"log"

	mathOperationsService "github.com/MitoVeli/math_grpc_server/pkg"
)

type MathOperationsServiceServer struct {
	UnimplementedMathOperationsServer
	mathOperationsService mathOperationsService.MathOperations
}

func (s *MathOperationsServiceServer) Add(ctx context.Context, in *OperationRequest) (*OperationResponse, error) {

	if err := s.mathOperationsService.Add(in.A, in.B, &in.Result); err != nil {
		log.Println("error in Add method", err)
		return nil, errors.New("failed to add")
	}

	// return user
	return &OperationResponse{
		Result: in.Result,
	}, nil
}
