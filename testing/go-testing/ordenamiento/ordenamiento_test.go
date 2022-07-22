package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input y output)
	numeros := []int{5, 2, 1, 3, 4}
	resultadoEsperado := []int{1, 2, 3, 4, 5}

	// Se ejecuta el test
	resultadoObtenido := Ordenar(numeros)

	// Se compara el resultado obtenido con el resultado esperado
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
}
