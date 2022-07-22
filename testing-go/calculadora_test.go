package testing_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 5

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "debe ser iguales")
}
