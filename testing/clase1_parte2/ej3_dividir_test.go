package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 0

	// Se ejecuta el test
	_, err := Dividir(num1, num2)

	// Se validan los resultados
	//assert.Equal(t, nil, err, err.Error())
	assert.Nil(t, err)

}
