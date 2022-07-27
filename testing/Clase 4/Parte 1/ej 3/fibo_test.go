package fibo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{4, 3},
		{5, 5},
		{14, 377},
	}
	for _, test := range tests {
		res := Fibonacci(test.given)
		assert.Equal(t, test.expected, res)
	}
}
