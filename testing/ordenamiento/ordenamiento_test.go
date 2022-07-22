package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderSlice(t *testing.T) {
	slice := []int{4, 32, 5, 2, 1}
	resultadoEsperado := []int{1, 2, 4, 5, 32}

	resultado := OrderSlice(slice)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
