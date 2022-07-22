package claseTesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	result := Restar(5, 3)

	assert.Equal(t, result, 2, "deben ser iguales")
}

func TestDividirSinDivisorCero(t *testing.T) {
	result, err := Dividir(6, 3)

	assert.Equal(t, result, 2.0, "deben ser iguales")
	assert.Nil(t, err, "no debe haber error")
}

func TestDividirConDivisorCero(t *testing.T) {
	_, err := Dividir(6, 0)

	assert.NotNil(t, err, "no debe haber error")
}
