package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	n1, n2 := 5, 0
	_, err := Dividir(n1, n2)
	assert.NotNil(t, err)
}
