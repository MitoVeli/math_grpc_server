package test

import (
	"testing"

	mathOperations "github.com/MitoVeli/math_grpc_server/pkg/math"
	"github.com/test-go/testify/assert"
)

func TestDoMath(t *testing.T) {

	firstNumber := float32(10)
	secondNumber := float32(5)
	invalidOperationSign := "&"
	var result float32

	mathOperationsService := mathOperations.NewMathOperationsService()

	t.Run("Add", func(t *testing.T) {

		err := mathOperationsService.Calculate(firstNumber, secondNumber, "+", &result)

		assert.Equal(t, result, firstNumber+secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Subtract", func(t *testing.T) {

		err := mathOperationsService.Calculate(firstNumber, secondNumber, "-", &result)

		assert.Equal(t, result, firstNumber-secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Multiply", func(t *testing.T) {

		err := mathOperationsService.Calculate(firstNumber, secondNumber, "*", &result)

		assert.Equal(t, result, firstNumber*secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Divide", func(t *testing.T) {

		err := mathOperationsService.Calculate(firstNumber, secondNumber, "/", &result)

		assert.Equal(t, result, firstNumber/secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Divide by zero", func(t *testing.T) {

		err := mathOperationsService.Calculate(0, 0, "/", &result)

		assert.Equal(t, "cannot divide by zero", err.Error())
	})

	t.Run("Invalid operation sign", func(t *testing.T) {

		err := mathOperationsService.Calculate(firstNumber, secondNumber, invalidOperationSign, &result)

		assert.Equal(t, "invalid operation sign: "+invalidOperationSign, err.Error())
	})
}
