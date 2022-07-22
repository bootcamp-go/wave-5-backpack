package clase1_parte2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	nums := []int{6, 4, 7, 3, 7, 3}
	resultadoEsperado := []int{3, 3, 4, 6, 7, 7}

	resultado := Ordenar(nums)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
