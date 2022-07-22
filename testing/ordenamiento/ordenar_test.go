package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	resultadoEsperado := []int{0, 1, 2, 3, 4, 5}

	nums := []int{4, 5, 0, 3, 2, 1}

	resultado := Ordenar(nums)

	assert.Equal(t, resultadoEsperado, resultado, "El reultado no es el esperado")
}
