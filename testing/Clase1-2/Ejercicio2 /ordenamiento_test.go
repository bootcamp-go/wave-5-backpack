package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrden(t *testing.T) {
	desorden := []int{2, 5, 3, 1, 4}
	esperado := []int{1, 2, 3, 4, 5}

	resultado := Ordenar(desorden)

	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}
