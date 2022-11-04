package pkg

type mathOperationsService struct {
}

func NewMathOperationsService() MathOperations {
	return &mathOperationsService{}
}

func (s *mathOperationsService) Add(x int32, y int32, result *int32) error {
	*result = x + y
	return nil
}

func (s *mathOperationsService) Subtract(x int32, y int32, result *int32) error {
	*result = x - y
	return nil
}

func (s *mathOperationsService) Multiply(x int32, y int32, result *int32) error {
	*result = x * y
	return nil
}

func (s *mathOperationsService) Divide(x int32, y int32, result *int32) error {
	*result = x / y
	return nil
}
