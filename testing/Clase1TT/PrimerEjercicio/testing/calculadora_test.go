package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1, num2, result := 3, 5, 8

	r := Restar(num1, num2)
	assert.NotNil(t, r)
	assert.Equal(t, r, result)
}
