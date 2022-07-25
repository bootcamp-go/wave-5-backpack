package C1TT

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	slice := []int{4, 3, 2, 1}
	resultadoEsperado := []int{1, 2, 3, 4}

	// Se ejecuta el test
	resultado := Ordenar(slice)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
