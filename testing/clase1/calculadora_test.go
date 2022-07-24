package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	num1 := 40
	num2 := 10
	HopeResult := 30
	result := Restar(num1, num2)

	// Se validan los resultados
	// if result != HopeResult {
	//     t.Errorf("Funcion resta() arrojo el resultado = %v, pero el esperado es  %v", result, HopeResult)
	// }

	assert.Equal(t, HopeResult, result, "deben ser iguales")
}
