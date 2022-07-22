package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRest(t *testing.T) {
	n1, n2 := 5, 4
	res := 1
	result := Resta(n1, n2)
	assert.Equal(t, res, result, "deben ser iguales")
}
