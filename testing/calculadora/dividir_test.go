package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T)  {
	num1 := 3
    num2 := 2
    

	_ , resultado := Dividir(num1, num2)

	assert.Nil(t, resultado)
}