package ejercicio2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	sliceEnteros := []int{9, 4, 2, 5, 3, 6}
	esperado := []int{2, 3, 4, 5, 6, 9}

	resultado := Ordenar(sliceEnteros)

	assert.Equal(t, esperado, resultado, "deben ser iguales")
}
