package ejercicio3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirCorrecto(t *testing.T) {
	num1 := 6
	num2 := 2
	resultadoEsperado := 3
	resultado, err := dividir(num1, num2)

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividirError(t *testing.T) {
	num1 := 3
	num2 := 0
	_, err := dividir(num1, num2)
	assert.EqualError(t, err, "error: el denominador no puede ser 0 o negativo")
}
