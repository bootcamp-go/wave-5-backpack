package clase1_parte2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 1
	num2 := 2
	resultadoEsperado := 1

	resultado := Restar(num1, num2)

	/*if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arrojó el resultado = %v, pero el esperado es = %v", resultado, resultadoEsperado)
	}*/

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
