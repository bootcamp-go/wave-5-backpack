package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	s := []int{15, 10, 5, 2, 13, 1, 6, 8}
	ordenado := []int{1, 2, 5, 6, 8, 10, 13, 15}

	assert.Equal(t, ordenado, Ordenar(s), "El slice obtenido es distinto al slice esperado")
}
