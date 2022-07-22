package ejercicio1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	num1 := 8
	num2 := 3
	esperado := 5

	resultado := Restar(num1, num2)

	assert.Equal(t, esperado, resultado, "deben ser iguales")
}
