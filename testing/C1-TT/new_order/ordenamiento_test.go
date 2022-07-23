package neworder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAscending(t *testing.T) {
	// Se inicializa el slice a ordenar
	newOrder := []int{12, 15, 13, 11, 10, 14}

	ascendResult := []int{10, 11, 12, 13, 14, 15}

	result := Ordenamiento(newOrder)

	assert.Equal(t, ascendResult, result, "Orden ascendente")
}
