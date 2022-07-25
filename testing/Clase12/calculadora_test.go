package Clase12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	var (
		num1     int64 = 9
		num2     int64 = 10
		esperado int64 = -1
	)
	res := Restar(num1, num2)
	assert.NotNil(t, res)
	assert.Equal(t, esperado, res)
}
