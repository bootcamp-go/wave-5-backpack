package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	numSlice := []int{3, 1, 4, 6, 2, 5, 7}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7}

	resultado := Ordenar(numSlice)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
