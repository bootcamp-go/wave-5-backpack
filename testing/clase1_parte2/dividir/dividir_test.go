package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirError(t *testing.T) {

	num1 := 8
	num2 := 0
	_, err := Dividir(num1, num2)
	assert.EqualError(t, err, "el denominador no puede ser 0")
}

func TestDividirExito(t *testing.T) {

	num1 := 8
	num2 := 1
	resultadoEsp := 8
	resultado, err := Dividir(num1, num2)

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsp, resultado, "deberían ser iguales")
}
