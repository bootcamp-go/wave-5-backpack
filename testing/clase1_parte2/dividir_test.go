package clase1_parte2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 2

	resultado, err := Dividir(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.Nil(t, err)
}
