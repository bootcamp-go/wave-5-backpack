package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 2
	resultadoEsperado := 3

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}
