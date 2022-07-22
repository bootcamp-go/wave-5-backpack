package calculadora

//Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	//Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8

	resultado := Sumar(num1, num2)

	//Test con libreria testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

	/* if resultado != resultadoEsperado {

		t.Errorf("Funcion suma() retorno el resultado = %v, pero el ePsperado es %v", resultado, resultadoEsperado)

	} */

}
func TestRestar(t *testing.T) {
	//Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 2
	resultadoEsperado := 3

	resultado := Restar(num1, num2)

	//Test con libreria testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

func TestDividir(t *testing.T) {

	num1 := 3
	num2 := 0

	resultadoEsperado := "El denominador no puede ser 0"

	resultado, err := Dividir(num1, num2)

	assert.Equal(t, resultadoEsperado, err.Error())
	assert.NotNil(t, resultado)
}
