package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 5
	resultadoEsperado := 6

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
*/

/*
func TestOrdenar(t *testing.T) {
	s := []int{4, 2, 3, 1}
	resultado := []int{1, 2, 4, 3}

	Ordenar(s)

	assert.Equal(t, s, resultado, "no estan ordenados")
}
*/

func TestDividir(t *testing.T) {
	n1 := 5
	n2 := 0

	_, err := Dividir(n1, n2)
	assert.Nil(t, err)

}
