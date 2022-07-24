package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	list := []int{5, 2, 7, 1}
	sortedList := []int{1, 2, 5, 7}
	ordenar(list)
	assert.Equal(t, sortedList, list, "deben ser iguales")
}
