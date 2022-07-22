package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	numbers := []int{5, 4, 7, 1, 3, 6, 2}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7}

	resultado := OrdenamientoAsc(numbers)

	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}
