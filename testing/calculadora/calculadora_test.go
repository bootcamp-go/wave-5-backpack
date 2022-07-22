package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 3
	num2 := 5
	resultadoEsperado := -2

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "El reultado no es el esperado")
}

func TestDividir(t *testing.T) {
	num1 := 3
	num2 := 0

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err, "Division por 0 no controlada")
}
