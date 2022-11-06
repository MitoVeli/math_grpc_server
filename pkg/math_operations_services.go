package pkg

type mathOperationsService struct {
}

func NewMathOperationsService() MathOperations {
	return &mathOperationsService{}
}

func (s *mathOperationsService) Add(x int64, y int64, result *int64) error {
	*result = x + y
	return nil
}

func (s *mathOperationsService) Subtract(x int64, y int64, result *int64) error {
	*result = x - y
	return nil
}

func (s *mathOperationsService) Multiply(x int64, y int64, result *int64) error {
	*result = x * y
	return nil
}

func (s *mathOperationsService) Divide(x int64, y int64, result *int64) error {
	*result = x / y
	return nil
}
