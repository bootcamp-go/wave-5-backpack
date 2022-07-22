package testing

import "github.com/stretchr/testify/assert"
import "testing"

func TestOrder(t *testing.T) {
	//Se inicializan los datos de input/output
	nums := []int{5, 3, 4, 7, 8, 9}
	expectedRes := []int{3, 4, 5, 7, 8, 9}

	//Ejecutamos el metodo
	res := Order(nums)

	//Se valida el resultado
	assert.Equal(t, expectedRes, res, "Deben ser iguales")
}
