package pkg

type MathOperations interface {
	Add(x int32, y int32, result *int32) error
	Subtract(x int32, y int32, result *int32) error
	Multiply(x int32, y int32, result *int32) error
	Divide(x int32, y int32, result *int32) error
}
