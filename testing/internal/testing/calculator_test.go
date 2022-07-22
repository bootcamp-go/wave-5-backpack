package testing

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

func TestRes(t *testing.T) {
	//Se inicializan los datos de input/output
	num1 := 1
	num2 := 1
	expectedRes := 0

	//Ejecutamos el metodo
	res := Res(num1, num2)

	//Se valida el resultado
	assert.Equal(t, expectedRes, res, "Deben ser iguales")
}

func TestDiv(t *testing.T) {
	//Se inicializan los datos de input/output
	num1 := 1
	num2 := 0

	//Ejecutamos el metodo
	res, err := Div(num1, num2)

	//Se valida el resultado
	if err != nil {
		assert.Nil(t, err, "El denominador no puede ser 0")
	}
	assert.NotNil(t, res, "El denominador no puede ser 0")
}
