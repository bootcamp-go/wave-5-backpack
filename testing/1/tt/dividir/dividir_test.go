package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivi(t *testing.T) {
	// Test resultado
	num := 10
	den := 2
	esperado := num / den

	obtenido, err := Dividir(num, den)

	assert.Equal(t, esperado, obtenido)
	assert.Nil(t, err)

	// Test error
	den = 0

	obtenido, err = Dividir(num, den)

	assert.Equal(t, 0, obtenido)
	assert.Error(t, err)
}
