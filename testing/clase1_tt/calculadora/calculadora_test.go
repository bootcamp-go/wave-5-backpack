package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
  num1 := 5
  num2 := 3
  esperado := 2

  res := Restar(num1, num2)

  assert.Equal(t, esperado, res)
}

func TestDividir(t *testing.T) {
	num1 := 10

	res, errGot := Dividir(num1, 0)
	assert.Equal(t, 0, res)
	assert.NotNil(t, errGot)
	assert.EqualError(t, errGot, "no se puede dividir por 0")

	res, err := Dividir(num1, 2)
	assert.Nil(t, err)
}
