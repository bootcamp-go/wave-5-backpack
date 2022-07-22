package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num := 10
	den1 := 1
	den2 := 0
	expectedResult := 5
	Result, err1 := Dividir(num, den1)
	_, err2 := Dividir(num, den2)

	assert.Equal(t, expectedResult, Result)
	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
