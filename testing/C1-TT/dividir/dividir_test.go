package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num := 8
	den := 0
	expResult := 2

	actual, err := Dividir(num, den)
	if err != nil {
		assert.EqualError(t, err, "El denominador no puede ser cero")
		return
	}
	assert.Equal(t, expResult, actual)
}
