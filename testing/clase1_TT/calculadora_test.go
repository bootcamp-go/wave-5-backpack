package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 3
	resultadoEsperado := 2
	resultado := Restar(num1, num2)
	// if resultadoEsperado !=resultado{
	// 	t.Errorf("Funcion Resta() arrojo el resultado = %v pero se esperaba %d ",resultado,resultadoEsperado)
	// }
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
func TestOrderAsc(t *testing.T) {
	sliceInt := []int{5, 1, 2, 7, 4, 8, 2, 9, 0}
	resultado := OrderAsc(sliceInt)

	for i, v := range resultado {
		if i == len(resultado)-1 {
			break
		}
		assert.GreaterOrEqual(t, resultado[i+1], v, "error on order slice[%d]= %d is bigger than slice[%d]= %d ", i, v, i+1, resultado[i+1])
		// if v < resultado[i+1] {
		// 	t.Errorf("error on order slice positon %d is bigger than %d", i, i+1)
		// }
	}
}

func TestDividir(t *testing.T) {
	den := 1
	num := 1

	// Se ejecuta el test funciona
	quo, err := Dividir(num, den)

	assert.Nil(t, err)
	assert.Equal(t, 1, quo)

	// Se ejecuta el test failure
	quoFail, err := Dividir(num, 0)

	// Se validan los resultados
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "the denominator should be different of 0")
		assert.Equal(t, 0, quoFail)
	}

}
