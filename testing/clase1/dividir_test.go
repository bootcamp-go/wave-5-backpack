package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 7
	num2 := 2
	hopeResult := 3
	result, err := dividir(num1, num2)
	if num2 == 0 {
		assert.NotNil(t, err, "se deveria tirar un error por ingresar un 0 de denominador")
	} else {
		assert.Nil(t, err, "el demoninador no fue 0 asi que no se deveria generar ningun error")
	}
	assert.Equal(t, hopeResult, result, "deben ser iguales")
}
