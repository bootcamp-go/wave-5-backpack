package testing_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDividir(t *testing.T) {
	_, err := Dividir(25, 0)
	assert.EqualError(t, err, "No se puede divir por 0, el denominador debe cambiarse")
}
