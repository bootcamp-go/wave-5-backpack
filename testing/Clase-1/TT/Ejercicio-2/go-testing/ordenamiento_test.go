package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {

	num := []int{2, 5, 1, -1, -20}
	resultadoEsperado := []int{-20, -1, 1, 2, 5}

	resultado := Ordenar(num)

	assert.Equal(t, resultadoEsperado, resultado, "Los slices deben ser iguales")
}
