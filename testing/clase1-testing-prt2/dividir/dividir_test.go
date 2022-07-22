package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestDividirTestify(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 3

	// Se ejecuta el test
	_, err := Dividir(num1, num2)

	// Se validan los resultados
	if err != nil {
		t.Fatal(err)
	}

	// Se validan los resultados aprovechando testify
	assert.Nil(t, err)

}
