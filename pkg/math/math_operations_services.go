package pkg

import (
	"errors"

	"github.com/MitoVeli/math_grpc_server/enums"
)

type mathOperationsService struct {
}

func NewMathOperationsService() MathOperations {
	return &mathOperationsService{}
}

func (s *mathOperationsService) DoMath(x float32, y float32, operationSign string, result *float32) error {

	switch operationSign {
	case enums.ADD:
		if err := s.add(x, y, result); err != nil {
			return err
		}
	case enums.SUBTRACT:
		if err := s.subtract(x, y, result); err != nil {
			return err
		}
	case enums.MULTIPLY:
		if err := s.multiply(x, y, result); err != nil {
			return err
		}
	case enums.DIVIDE:
		if y == 0 {
			return errors.New("cannot divide by zero")
		}
		if err := s.divide(x, y, result); err != nil {
			return err
		}
	default:
		return errors.New("invalid operation sign: " + operationSign)
	}

	return nil
}

func (s *mathOperationsService) add(x float32, y float32, result *float32) error {
	*result = x + y
	return nil
}

func (s *mathOperationsService) subtract(x float32, y float32, result *float32) error {
	*result = x - y
	return nil
}

func (s *mathOperationsService) multiply(x float32, y float32, result *float32) error {
	*result = x * y
	return nil
}

func (s *mathOperationsService) divide(x float32, y float32, result *float32) error {
	*result = x / y
	return nil
}
