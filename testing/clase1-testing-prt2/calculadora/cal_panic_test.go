package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 1

	// Se ejecuta el test
	resultado := Dividir(num1, num2)

	// Se validan los resultados aprovechando testify
	assert.NotNil(t, resultado)

}
