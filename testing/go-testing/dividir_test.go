package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	_, err := Dividir(5, 0)
	assert.EqualError(t, err, "El denominador no puede ser 0")
}
