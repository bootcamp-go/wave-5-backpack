package dividir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDividir(t *testing.T) {
	num1 := 35
	num2 := 5
	resultadoEsperado := 7

	resultado, err := Dividir(num1, num2)
	if num2 == 0 {
		assert.NotNil(t, err)
	} else {
		assert.Nil(t, err)
		assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	}

}
