package pkg

type MathOperations interface {
	Add(x int64, y int64, result *int64) error
	Subtract(x int64, y int64, result *int64) error
	Multiply(x int64, y int64, result *int64) error
	Divide(x int64, y int64, result *int64) error
}
