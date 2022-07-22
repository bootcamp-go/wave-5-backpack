package dividir

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirErr(t *testing.T) {
	// Se inicializan los datos a usar en el test (input y output)
	num1 := 5
	num2 := 0
	errorEsperado := errors.New(ErrDivisionCero)

	// Se ejecuta el test
	_, err := Dividir(num1, num2)

	// Se compara el resultado obtenido con el resultado esperado
	assert.Equal(t, errorEsperado, err)
}

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input y output)
	num1 := 6
	num2 := 2
	resultadoEsperado := 3

	// Se ejecuta el test
	resultadoObtenido, err := Dividir(num1, num2)

	// Se compara el resultado obtenido con el resultado esperado
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.Nil(t, err)
}
