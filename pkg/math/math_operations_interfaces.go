package pkg

type MathOperations interface {
	DoMath(x float32, y float32, operationSign string, result *float32) error
}
