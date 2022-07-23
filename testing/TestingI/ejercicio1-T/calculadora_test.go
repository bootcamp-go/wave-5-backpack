package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 4
	num2 := 4
	expected := 0

	res := Restar(num1, num2)

	assert.Equal(t, expected, res, "Resultado diferente al esperado")
}
