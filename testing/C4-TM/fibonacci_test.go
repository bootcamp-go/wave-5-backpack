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
		{3, 2},
		{7, 13},
		{12, 144},
		{17, 1597},
	}

	for _, d := range tests {
		got := fibonacci(d.arg)
		assert.Equal(t, d.want, got)
	}
}
