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

func (s *mathOperationsService) DoMath(x int64, y int64, operationSign string, result *int64) error {

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
		return errors.New("invalid operation sign, operation sign:" + operationSign)
	}

	return nil
}

func (s *mathOperationsService) add(x int64, y int64, result *int64) error {
	*result = x + y
	return nil
}

func (s *mathOperationsService) subtract(x int64, y int64, result *int64) error {
	*result = x - y
	return nil
}

func (s *mathOperationsService) multiply(x int64, y int64, result *int64) error {
	*result = x * y
	return nil
}

func (s *mathOperationsService) divide(x int64, y int64, result *int64) error {
	*result = x / y
	return nil
}
