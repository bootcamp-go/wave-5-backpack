package C1_P2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdenamiento(t *testing.T) {
	cases := struct {
		toOrder []int
		sort    []int
	}{
		toOrder: []int{1, 4, 2, 1, 8, 9},
		sort:    []int{1, 1, 2, 4, 8, 9},
	}
	expected := Ordenamiento(cases.toOrder)
	assert.Equal(t, cases.sort, expected)
}
