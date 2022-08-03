package calculadora

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

	resultado := Restar(num1, num2)
	assert.NotNil(t, resultado)
	assert.Equal(t, esperado, resultado)
}

// Test function Restar using the standard library
/* func TestRestar(t *testing.T) {
	var (
		num1   int64 = 9
		num2   int64 = 10
		result int64 = -1
	)

	if got := Restar(num1, num2); got != result {
		t.Errorf("Failed, expected %d but got %d", result, got)
	}
} */

// Test function Restar using the standard library with Tables
/* func TestRestar(t *testing.T) {
	arrange := []struct {
		num1   int64
		num2   int64
		result int64
	}{
		{10, 9, 1},
		{7, 9, -2},
		{5, 6, -1},
		{0, 9, -9},
	}

	for _, value := range arrange {
		if got := Restar(value.num1, value.num2); got != value.result {
			t.Errorf("Failed, expected %d but got %d", value.result, got)
		}
	}
} */
