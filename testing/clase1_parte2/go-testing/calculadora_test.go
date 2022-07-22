package go_testing

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestRestar(t *testing.T) {
	//Inicializacion de input/output
	num1 := 5
	num2 := 3
	expectedResult := 2

	//Ejecuto test
	result := Restar(num1, num2)

	//Valido resultados
	assert.Equal(t, expectedResult, result, "not are equals")
}


func TestDividir(t *testing.T) {
	// -------Dividir OK-------
	num1 := 9
	num2 := 3
	expectedResult := 3

	//Ejecuto test
	result, err := Dividir(num1, num2)
	//Valido resultados
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result, "not are equals")



	// -------Dividir error-------
	num2 = 0
	result, err = Dividir(num1, num2)

	assert.Error(t, err)
	assert.Equal(t, errors.New("the denominator number should not be equal to 0"), err, "not equals")
	assert.NotEqual(t, expectedResult, result, "are equals")

}


