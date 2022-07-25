package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	n1, n2 := 3, 2
	r := 1
	result := Restar(n1, n2)
	assert.Equal(t, r, result, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	n1, n2 := 5, 0
	_, err := Dividir(n1, n2)
	assert.NotNil(t, err)
}

func TestOrdenamiento(t *testing.T) {
	s := make([]int, 3)
	s[0], s[1], s[2] = 3, 2, 1

	res := make([]int, 3)
	res[0], res[1], res[2] = 1, 2, 3
	result := OrdenarSlice(s)
	assert.Equal(t, res, result, "deben ser iguales")
}
