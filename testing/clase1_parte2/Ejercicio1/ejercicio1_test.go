package ejercicio1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 3
	num2 := 5
	resultadoEsperado := 2

	assert.Equal(t, resultadoEsperado, restar(num1, num2), "deben ser iguales")
}
