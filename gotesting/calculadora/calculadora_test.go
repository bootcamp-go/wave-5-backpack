package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// inicializamos los valores
	num1 := 8
	num2 := 5
	resultadoEsperado := 3

	//ejecutamos el test
	resultadoReal := Restar(num1, num2)

	// validamos
	assert.Equal(t, resultadoEsperado, resultadoReal, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	// inicializamos los valores
	num := 8
	den := 2
	resultadoEsperado := 4

	//ejecutamos el test
	resultadoReal, _ := Dividir(num, den)

	// validamos
	assert.Equal(t, resultadoEsperado, resultadoReal, "deben ser iguales")

}

func TestDiviriDenCero(t *testing.T) {
	// inicializamos los valores
	num := 8
	den := 0

	//ejecutamos el test
	_, err := Dividir(num, den)

	// validamos que el error no sea nil
	assert.NotNil(t, err)
}
