package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestOrdenamiento(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	numerosDesordenados := []int{5, 3, 1}
	resultadoEsperado := []int{1, 3, 5}

	// Se ejecuta el test
	resultado := OrdernarEnteros(numerosDesordenados)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
