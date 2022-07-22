package clase_1_parte_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	numerosSlice := []int{5, 4, 3, 7, 8, 2, 1}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 7, 8}

	resultado := Ordenar(numerosSlice)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
