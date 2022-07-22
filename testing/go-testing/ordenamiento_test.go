package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAscendingSort(t *testing.T) {
	unsortedList := []int{0, -2, 54, 21, 75, 12}
	sortedList := []int{-2, 0, 12, 21, 54, 75}

	result := AscendingSort(unsortedList)

	assert.Equal(t, sortedList, result, "deben ser iguales")
}
