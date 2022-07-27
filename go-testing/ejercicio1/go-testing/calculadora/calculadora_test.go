package calculadora

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	var (
		num1     int = 8
		num2     int = 5
		esperado int = 3
	)

	resultado := Restar(num1, num2)
	assert.NotNil(t, resultado)
	assert.Equal(t, esperado, resultado)
}
