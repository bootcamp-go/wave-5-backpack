package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{1, 0}, {2, 1}, {3, 1}, {4, 2}, {5, 3}, {6, 5}, {7, 8}}
	for _, test := range tests {
		result := fibonacci(test.arg)
		assert.Equal(t, test.want, result)
	}
}
