package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonaci(t *testing.T) {
	type testFibonaci struct {
		input     int
		fibonacci int
	}
	// 1  2  3  4  5  6  7   8   9   10  11  12   13   14   15   16
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987
	data := []testFibonaci{{1, 1}, {5, 5}, {10, 55}, {15, 610}}

	for _, e := range data {
		output := fibonacci(e.input)
		assert.Equal(t, e.fibonacci, output)
	}
}
