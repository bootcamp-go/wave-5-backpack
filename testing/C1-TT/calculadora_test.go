package c1tt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 3
	resultadoEsperado := 2
	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividirPorCero(t *testing.T) {
	num1 := 5
	num2 := 0

	errorEsperado := errors.New("el denominador no puede ser 0")

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err, "debe retornar un error con denominador 0")
	assert.Equal(t, errorEsperado, err, "el mensaje de error es incorrecto")

}
