package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	caso := struct {
		ordenar  []int
		ordenado []int
	}{
		ordenar:  []int{5,6,2,3,4,7},
		ordenado: []int{2,3,4,5,6,7},
	}

	resultado := Ordenamiento(caso.ordenar)
	assert.Equal(t, caso.ordenado, resultado)
}
