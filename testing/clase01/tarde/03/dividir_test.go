package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDividirZero(t *testing.T){
// 	num := 6
// 	den := 0

// 	_, err := Dividir(num, den)

// 	assert.NotNil(t, err)
// }

func TestDividir(t *testing.T){
	num := 6
	den := 2

	resultadoEsperado :=3

	resultado,err := Dividir(num, den)
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado)
}