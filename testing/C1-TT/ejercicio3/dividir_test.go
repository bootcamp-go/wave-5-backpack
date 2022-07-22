package ejercicio3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num := 22
	den := 11

	esperado := 3

	resultado, err := Dividir(num, den)

	assert.Nil(t, err, err)
	assert.Equal(t, esperado, resultado, "debe ser igual")
}
