package calculadora

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T)  {
	list := []int{2, 1, 6, 3, 7}
	resultadoEsperado := []int{1, 2, 3, 6, 7}

	resultado := Ordenamiento(list)

	assert.Equal(t, resultadoEsperado, resultado, "No se encuentra ordenado el Slice")
}
