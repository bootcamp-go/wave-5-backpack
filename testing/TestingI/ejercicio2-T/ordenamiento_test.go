package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {

	list := []int{3, 2, 5, 4, 1}

	listExpected := []int{1, 2, 3, 4, 5}
	res := Ordenar(list)

	assert.Equal(t, listExpected, res, "error al ordenar")
}
