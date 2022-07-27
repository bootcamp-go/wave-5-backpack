package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Aplicando TDD con Fibonacci

func TestFibonnac(t *testing.T) {
	n := 1
	res := Fibonacci(n)

	if n == 0 {
		assert.Equal(t, 0, res)
	} else if n == 1 {
		assert.Equal(t, 1, res)
	} else if n > 1 {

	}
}
