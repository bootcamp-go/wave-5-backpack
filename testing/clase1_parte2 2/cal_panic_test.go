package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisionFail(t *testing.T) {
	num1 := 4
	num2 := 0

	resultado := Dividir(num1, num2)

	assert.NotNil(t, nil, resultado)
}
