package ordenamiento

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestOrdenamiento ( t *testing.T){
	miSlice := []int{5,3,4,1,2}

	resultadoEsperado := []int{1,2,3,4,5}
	resultado := Ordenar(miSlice)

	assert.Equal(t,resultadoEsperado,resultado, "hubo un error al ordenar")
}