package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 3
	num2 := 1
	expectedResult := 2

	result := Restar(num1, num2)

	assert.Equal(t, expectedResult, result, "deben ser iguales")
}
