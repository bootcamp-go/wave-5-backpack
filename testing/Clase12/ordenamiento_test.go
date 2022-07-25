package Clase12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	caso := struct {
		ordenar  []int
		ordenado []int
	}{
		ordenar:  []int{1, 2, 3, 9, 2, 1},
		ordenado: []int{1, 1, 2, 2, 3, 9},
	}

	res := Ordenamiento(caso.ordenar)
	assert.Equal(t, caso.ordenado, res)
}
