package fibonaccitest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	test := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1}}

	for _, d := range test {
		got, _ := Fibonacci(d.arg)
		assert.Equal(t, d.want, got)
	}
	_, err := Fibonacci(-5)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "el numero debe ser mayor a 0")
}
