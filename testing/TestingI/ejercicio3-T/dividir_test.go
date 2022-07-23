package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 9
	num2 := 3

	expected := 3

	res, _ := Dividir(num1, num2)

	assert.Equal(t, expected, res, "no se completo la division")
}

func TestDividirCero(t *testing.T) {

	num1 := 2
	num2 := 0

	_, err := Dividir(num1, num2)
	assert.NotNil(t, err)
}
