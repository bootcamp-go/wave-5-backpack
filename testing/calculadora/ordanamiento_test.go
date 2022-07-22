package calculadora

import (
    "testing"

    "github.com/stretchr/testify/assert" 
)

func TestOrdenar(t *testing.T)  {
	num1 := []int{1, 3, 6, 2, 5, 4}
    resultadoEsperado := []int{1, 2, 3, 4, 5, 6}

	resultado := Ordenar(num1)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
