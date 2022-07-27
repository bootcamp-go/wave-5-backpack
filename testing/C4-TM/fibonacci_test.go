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
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{7, 13},
		{12, 144},
		{17, 1597},
		{23, 28657},
	}

	for _, d := range tests {
		got, _ := fibonacci(d.arg)
		assert.Equal(t, d.want, got)
	}

	_, err := fibonacci(-5)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "n debe ser >= 0")
}
