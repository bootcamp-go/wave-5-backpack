package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input y output)
	num1 := 5
	num2 := 2
	resultadoEsperado := 3

	// Se ejecuta el test
	resultadoObtenido := Restar(num1, num2)

	// Se compara el resultado obtenido con el resultado esperado
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
}
