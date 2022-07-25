package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {

	num := 2
	den := 1
	resultadoEsperado := 2

	resultado, err := Dividir(num, den)

	assert.Nil(t, err)

	assert.Equal(t, resultadoEsperado, resultado, "La división no se realizó de manera correcta")

}
