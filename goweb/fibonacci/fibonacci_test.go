package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	resultExpected1 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597}
	result1 := fibonacci(18)

	resultExpected2 := []int{0, 1, 1}
	result2 := fibonacci(3)

	resultExpected3 := []int{0, 1}
	result3 := fibonacci(1)

	assert.Equal(t, resultExpected1, result1)
	assert.Equal(t, resultExpected2, result2)
	assert.Equal(t, resultExpected3, result3)
}
