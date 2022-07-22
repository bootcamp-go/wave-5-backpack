package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	s := make([]int, 3)
	s[0], s[1], s[2] = 3, 2, 1

	res := make([]int, 3)
	res[0], res[1], res[2] = 1, 2, 3
	result := OrdenarSlice(s)
	assert.Equal(t, res, result, "deben ser iguales")
}
