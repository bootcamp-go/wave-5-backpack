package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {

	num1 := 8
	num2 := 0
	_, err := Dividir(num1, num2)

	assert.Nil(t, err)
}
