package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{1, 1},
		{2, 1},
	}

	for _, d := range tests {
		got := fibonacci(d.arg)
		assert.Equal(t, d.want, got)
	}
}
