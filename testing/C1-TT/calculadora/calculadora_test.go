package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	//Se inicializan los datos a usar en el test (input/output)
	num1 := 6
	num2 := 4
	expectedResult := 2

	//Se ejecuta el test
	resultado := Restar(num1, num2)

	//Se validad los resultados
	assert.Equal(t, expectedResult, resultado, "deben ser iguales")
}
