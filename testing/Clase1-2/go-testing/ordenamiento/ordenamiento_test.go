package ordenamiento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdenamiento(t *testing.T) {
	sliceDesordenado := []int{8, 7, 3, 1, 9, 0}
	sliceOrdenado := []int{0, 1, 3, 7, 8, 9}

	resultado := ordenamientoBurbuja(sliceDesordenado)
	assert.Equal(t, resultado, sliceOrdenado, "deben ser iguales")
}
