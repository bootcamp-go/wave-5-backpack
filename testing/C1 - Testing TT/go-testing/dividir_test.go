package ejercicio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 4
	num2 := 2
	num3 := 0

	// Se ejecuta el test funciona
	resultadoWork, err := Dividir(num1, num2)

	assert.Nil(t, err)
	assert.Equal(t, 2, resultadoWork)

	// Se ejecuta el test failure
	resultadoFailure, err := Dividir(num1, num3)

	// Se validan los resultados aprovechando testify
	assert.NotNil(t, err)
	assert.EqualError(t, err, "El denominador no puede ser 0")
	assert.Equal(t, 0, resultadoFailure)
}
