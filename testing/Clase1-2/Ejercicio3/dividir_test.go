package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	a := 3
	b := 5
	esperado := a / b

	resultad, err := Dividir(a, b)

	assert.Nil(t, err)

	assert.Equal(t, esperado, resultad, "No son iguales")

}
