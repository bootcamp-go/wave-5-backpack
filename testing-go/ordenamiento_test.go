package testing_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdenAscendente(t *testing.T) {
	listaEnteros := []int{3, 8, 1, 9, 56, 78, 23, 90}
	enterosOrdenados := []int{1, 3, 8, 9, 23, 56, 78, 90}

	resultado := OrdenAscendente(listaEnteros)

	assert.Equal(t, enterosOrdenados, resultado, "Orden debe ser identico")
}