package ordenador

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultado := Ordenar([]int{2, 10, 1, 3, 7, 4, 8, 5, 9, 6})

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestOrdenarMultiple(t *testing.T) {
	var tests = []struct {
		lista []int
		want  []int
		err   string
	}{
		{lista: []int{2, 10, 1, 3, 7, 4, 8, 5, 9, 6}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{lista: []int{2, -10, 10, 3, -7, 4, 8, 5, 9, 6}, want: []int{-10, -7, 2, 3, 4, 5, 6, 8, 9, 10}},
		{lista: []int{0, -1, 1}, want: []int{-1, 0, 1}},
		{lista: []int{20, 10, 30, -10, -20, -30, 0}, want: []int{-30, -20, -10, 0, 10, 20, 30}},
		{lista: []int{1, 2, 3, 4, 5, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		{lista: []int{-1, 6, -1, 0, 6, 1}, want: []int{-1, -1, 0, 1, 6, 6}},
	}
	for i, tt := range tests {
		testname := fmt.Sprintf("ordenar %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := Ordenar(tt.lista)
			assert.Equal(t, tt.want, ans, "deben ser iguales")
		})
	}
}
