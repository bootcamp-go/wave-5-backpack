package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestOrdnamientoTestify(t *testing.T) {
	// Se inicializan el Slice de datos (int) resultado esperado
	listResult := []int{9, 8, 7, 5, 4, 3}

	// Se ejecuta el test
	resultado := Ordenamiento([]int{5, 3, 4, 7, 8, 9})

	// Se validan los resultados aprovechando testify
	assert.Equal(t, listResult, resultado, "El resultado esperado: %v", listResult)

}
