package restar

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestRestarTestify(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -2

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}

	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "El resultado esperado %v", resultadoEsperado)
}
