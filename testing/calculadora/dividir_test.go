package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num := 15
	den := 0
	resultadoEsperado := 3

	resultado, err := Dividir(num, den)
	assert.Nil(t, err, "El denominador no puede ser cero")
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}
