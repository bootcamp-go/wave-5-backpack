package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	n1 := 10
	n2 := 5
	resWished := 5

	res := Restar(n1, n2)

	assert.Equal(t, resWished, res, "Deberían ser iguales los resultados")

}

func TestDividir(t *testing.T) {
	numerador := 10
	denominador := 0
	resWished := 2

	res, err := Dividir(numerador, denominador)

	if denominador == 0 {
		assert.Nil(t, err.Error(), "El dividendo no puede ser 0")
	} else {
		assert.Equal(t, resWished, res, "Deberían ser iguales los resultados")

	}
}

func TestSort(t *testing.T) {
	values := []int{1, 4, 3, 2}
	res := []int{1, 2, 3, 4}

	s := Sort(values...)

	assert.Equal(t, res, s, "Los arreglos deberían estar ordenados de manera ascendente")

}
