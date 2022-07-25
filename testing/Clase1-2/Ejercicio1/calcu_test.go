package calcu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	a := 1
	b := 2
	r := 3

	r2 := suma(a, b)

	assert.Equal(t, r, r2, "Deben ser iguales")
}

//Ejercicio 1
func TestResta(t *testing.T) {
	a := 3
	b := 2
	r := 1

	r2 := resta(a, b)

	assert.Equal(t, r, r2, "Deben ser iguales")
}
