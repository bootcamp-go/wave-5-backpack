package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 3
	num2 := 4

	resultadoEsperado := 8

	resultado := Sumar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	num1 := 2
	num2 := 0

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err, "no se puede dividir por cero")
}
