package Dividr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDividir(t *testing.T) {
	_, err := Dividir(50, 0)

	//usando testing
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//usando testify
	assert.Nil(t, err)
}
