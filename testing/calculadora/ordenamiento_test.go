package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	num := []int{5, 1, 2, 3, 45}
	resultadoEsperado := []int{1, 2, 3, 5, 45}

	resultado := Quicksort(num)

	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}
