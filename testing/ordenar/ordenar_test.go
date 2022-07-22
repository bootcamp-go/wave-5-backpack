package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarAsc(t *testing.T) {

	var sliceNum = []int{2, 3, 4, 5, 1}
	resultadoEsperado := []int{1, 2, 3, 4, 5}

	resultadoReal := OrdenasAsc(sliceNum)

	assert.Equal(t, resultadoEsperado, resultadoReal, "deben ser iguales")
}
