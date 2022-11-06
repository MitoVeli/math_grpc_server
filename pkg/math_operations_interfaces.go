package pkg

type MathOperations interface {
	DoMath(x int64, y int64, operationSign string, result *int64) error
}
