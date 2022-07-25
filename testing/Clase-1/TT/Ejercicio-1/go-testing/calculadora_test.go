package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRe(t *testing.T) {

	num1 := 4
	num2 := 2
	resultadoEsperado := 1

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Las restas deben ser iguales")
}
