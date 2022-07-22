package calculator

import "github.com/stretchr/testify/assert"
import "testing"

func TestSum(t *testing.T) {
	//Se inicializan los datos de input/output
	num1 := 1
	num2 := 1
	expectedRes := 2

	//Ejecutamos el metodo
	res := Sum(num1, num2)

	//Se valida el resultado
	assert.Equal(t, expectedRes, res, "Deben ser iguales")
}
