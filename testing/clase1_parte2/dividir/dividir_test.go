package dividir

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Table drive test para multiples casos
var dividirTests = []struct {
	testName                    string
	num, den, resultadoEsperado int
	errorEsperado               error
}{
	{"Test1", 10, 2, 5, nil},
	{"Test2", 5, 0, 0, fmt.Errorf("el denominador no puede ser 0")},
}

func TestDividir(t *testing.T) {
	for _, test := range dividirTests {
		t.Run(test.testName, func(t *testing.T) {
			num := test.num
			den := test.den
			resultadoEsperado := test.resultadoEsperado
			errorEsperado := test.errorEsperado

			resultado, err := Dividir(num, den)
			assert.Equal(t, errorEsperado, err, "deben ser iguales")
			assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

		})
	}
}
