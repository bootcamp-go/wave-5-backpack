package Ordenar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdenamiento(t *testing.T) {
	resultadoEsperado := []int{3, 4, 5, 6, 7, 8, 9, 10}

	resultado := Ordenamiento([]int{10, 5, 3, 4, 6, 7, 8, 9})

	assert.Equal(t, resultadoEsperado, resultado, "El resultado esperado es %v", resultadoEsperado)
}
