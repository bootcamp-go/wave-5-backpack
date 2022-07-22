package go_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	numbers := []int{3, 5, 1, 2, 6, 4}
	expectedNumbers := []int{1, 2, 3, 4, 5, 6}

	result := Ordenar(numbers)
	assert.NotEmpty(t, result)
	assert.Equal(t, expectedNumbers, result, "slices not equals")
}
