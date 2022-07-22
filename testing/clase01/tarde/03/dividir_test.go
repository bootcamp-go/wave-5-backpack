package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirZero(t *testing.T){
	num := 6
	den := 0

	_, err := Dividir(num, den)

	assert.NotNil(t, err)
}

func TestDividir(t *testing.T){
	num := 6
	den := 2

	resultadoEsperado :=3

	resultado,_ := Dividir(num, den)

	assert.Equal(t, resultadoEsperado, resultado)
}