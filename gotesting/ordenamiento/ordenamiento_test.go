package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarAsc(t *testing.T) {
	// declaramos
	var sliceNum = []int{5, 3, 4, 1, 2}
	resultadoEsperado := []int{1, 2, 3, 4, 5}

	// ejecutamos el test
	resultadoReal := OrdenasAsc(sliceNum)

	// validamos
	assert.Equal(t, resultadoEsperado, resultadoReal, "deben ser iguales")
}
