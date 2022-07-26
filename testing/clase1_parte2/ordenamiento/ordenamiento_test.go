package ordenamiento

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T)  {
	numeros := []int{1, 5, 4, 9, 7, 3}
	resultadoEsperado := []int{1, 3, 4, 5, 7, 9}

	resultado := Ordenar(numeros)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}