package calculadora

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestDividir(t *testing.T)  {

	num1 := 9
	num2 := 3
	num3 := 0

	resultadoOK, err := Dividir(num1, num2)

	assert.Nil(t, err)
	assert.Equal(t, 3, resultadoOK)


	resultadoFail, err := Dividir(num1, num3)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "No se puede dividir por 0")
	assert.Equal(t, 0, resultadoFail)


}
