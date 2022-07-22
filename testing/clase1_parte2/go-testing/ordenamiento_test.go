package claseTesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	arrayOrdenado := Ordernar([]int{5, 2, 9, 4, 29, 54, 0})

	assert.Equal(t, arrayOrdenado, []int{0, 2, 4, 5, 9, 29, 54}, "deben ser iguales")
}
