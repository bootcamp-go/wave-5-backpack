package dividir

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Tests = []struct {
	testName                    string
	num, den, resultadoEsperado int
	errorEsperado               error
}{
	{"correct_den", 20, 10, 3, nil},
	{"incorrect_den", 20, 0, 0, errors.New("el denominador no puede ser 0")},
}

func TestDividir(t *testing.T) {
	for _, test := range Tests {
		t.Run(test.testName, func(t *testing.T) {
			num := test.num
			den := test.den
			resultadoEsperado := test.resultadoEsperado
			errorEsperado := test.errorEsperado

			resultado, err := Dividir(num, den)

			assert.Equal(t, errorEsperado, err, "Deben ser iguales")
			assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
		})
	}
}
