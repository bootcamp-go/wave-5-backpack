package ejercicio

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestarSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 3
	resultadoEsperadoResta := 2
	resultadoEsperadoSuma := 8

	// Se ejecuta el test
	resultadoResta := Restar(num1, num2)

	// Se ejecuta el test
	resultadoSuma := Sumar(num1, num2)

	// Se validan los resultados de la resta
	assert.Equal(t, resultadoEsperadoResta, resultadoResta, "El valor esperado de la resta no coincide con el valor obtenido en la función")

	// Se validan los resultados de la suma
	assert.Equal(t, resultadoEsperadoSuma, resultadoSuma, "El valor esperado de la suma no coincide con el valor obtenido en la función")

}
