package dividir

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestDividir(t *testing.T)  {
	division, err := Dividir(25, 0)
	assert.Equal(t, 0, division)
	assert.Equal(t, errors.New("El denominador no puede ser 0"), err)
}