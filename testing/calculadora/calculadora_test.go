package calculadora

import (
    "testing"

    "github.com/stretchr/testify/assert" 
)

func TestResta(t *testing.T)  {
	num1 := 3
    num2 := 5
    resultadoEsperado := -2

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
