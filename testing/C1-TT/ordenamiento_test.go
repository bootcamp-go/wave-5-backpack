package c1tt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	numeros := []int{1, 8, 7, 5, 4, 3, 9, 2, 6}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultado := Ordenar(numeros)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
