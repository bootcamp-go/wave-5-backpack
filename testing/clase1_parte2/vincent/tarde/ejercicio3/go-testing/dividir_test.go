package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	var (
		num int = 10
		dem int = 2
	)

	resultado, err := Dividir(num, dem)
	assert.Nil(t, err)

	esperado := num / dem
	assert.Equal(t, esperado, resultado)
}
