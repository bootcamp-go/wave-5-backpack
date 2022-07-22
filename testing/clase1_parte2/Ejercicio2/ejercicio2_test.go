package ejercicio2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	sliceInt := []int{3, 6, 7, 1, 128, 63}
	resultadoEsperado := []int{1, 3, 6, 7, 63, 128}

	assert.Equal(t, resultadoEsperado, ordernar(sliceInt), "deben ser iguales")
}
