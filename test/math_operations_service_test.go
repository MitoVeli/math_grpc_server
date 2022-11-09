package test

import (
	"testing"

	mathOperations "github.com/MitoVeli/math_grpc_server/pkg/math"
	"github.com/test-go/testify/assert"
)

func TestDoMath(t *testing.T) {

	firstNumber := float32(10)
	secondNumber := float32(5)
	var result float32

	mathOperationsService := mathOperations.NewMathOperationsService()

	t.Run("Add", func(t *testing.T) {

		err := mathOperationsService.DoMath(firstNumber, secondNumber, "+", &result)

		assert.Equal(t, result, firstNumber+secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Subtract", func(t *testing.T) {

		err := mathOperationsService.DoMath(firstNumber, secondNumber, "-", &result)

		assert.Equal(t, result, firstNumber-secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Multiply", func(t *testing.T) {

		err := mathOperationsService.DoMath(firstNumber, secondNumber, "*", &result)

		assert.Equal(t, result, firstNumber*secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Divide", func(t *testing.T) {

		err := mathOperationsService.DoMath(firstNumber, secondNumber, "/", &result)

		assert.Equal(t, result, firstNumber/secondNumber)
		assert.Nil(t, err)
	})

	t.Run("Divide by zero", func(t *testing.T) {

		err := mathOperationsService.DoMath(0, 0, "/", &result)

		assert.Equal(t, "cannot divide by zero", err.Error())
	})

	t.Run("Bad operation operator", func(t *testing.T) {

		err := mathOperationsService.DoMath(firstNumber, secondNumber, "&", &result)

		assert.Equal(t, "invalid operation sign, operation sign:&", err.Error())
	})
}
