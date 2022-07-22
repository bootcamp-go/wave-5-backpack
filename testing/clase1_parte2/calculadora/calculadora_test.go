package calculadora

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Sumar(num1, num2)

	//  Se validan los resultados
	/* 	if resultado != resultadoEsperado {
		t.Errorf("Funcion sumar() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	} */

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 6
	num2 := 2
	resultadoEsperado := 4

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	//  Se validan los resultados
	/* 	if resultado != resultadoEsperado {
		t.Errorf("Funcion restar() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	} */
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
