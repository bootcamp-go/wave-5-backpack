package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	num1 := 10
	num2 := 5
	resEsperado := 5
	res := Restar(num1, num2)

	assert.Equal(t, res, resEsperado, "deberían ser iguales")
}
