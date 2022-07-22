package clase_1_parte_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirError(t *testing.T) {
	num := 3
	den := 0

	_, err := Dividir(num, den)

	assert.NotNil(t, err, "no hay error cuando denominador es 0")
	assert.EqualValues(t, ErrDominadorCero, err, "el error enviado no es el indicado")

}
