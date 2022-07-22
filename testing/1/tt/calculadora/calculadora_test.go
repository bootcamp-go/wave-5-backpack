package calculadora

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := rand.Int()
	num2 := rand.Int()
	resta := num1 - num2

	assert.Equal(t, resta, Restar(num1, num2), "El valor obtenido es distinto al valor esperado")
}
