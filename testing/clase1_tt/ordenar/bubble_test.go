package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{4, 3, 2, 8, 4, 5, 7, 9, 3, 21, 0}
	esperado := []int{0, 2, 3, 3, 4, 4, 5, 7, 8, 9, 21}

	resultado := bubbleSort(nums, len(nums))

	assert.Equal(t, esperado, resultado)
}
