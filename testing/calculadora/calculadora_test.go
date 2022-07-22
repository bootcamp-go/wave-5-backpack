package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 2
	num2 := 5
	resultadoEsperado := -3

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}
