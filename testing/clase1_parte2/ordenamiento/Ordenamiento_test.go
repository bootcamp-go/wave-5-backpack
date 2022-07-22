package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {

	sliceInt := []int{4, 2, 3, 656, 3, 123, 7, 3, 4}
	resultadoEsperado := []int{2, 3, 3, 3, 4, 4, 7, 123, 656}
	resultado := Ordenar(sliceInt)

	assert.Equal(t, resultado, resultadoEsperado, "deberían ser iguales")
}
