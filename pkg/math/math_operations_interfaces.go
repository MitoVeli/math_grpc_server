package pkg

type MathOperations interface {
	DoMath(x float64, y float64, operationSign string, result *float64) error
}
