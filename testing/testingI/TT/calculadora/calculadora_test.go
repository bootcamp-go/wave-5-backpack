package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 20
	num2 := 15
	resultadoEsperado := 5

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("funcion Resta() arrojo el resultado = %v, pero el resultado esperado es %v.", resultado, resultadoEsperado)
	}
}

func TestRestarMultiple(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, -1},
		{1, 0, 1},
		{2, -2, 4},
		{0, -1, 1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Restar(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestDividir(t *testing.T) {
	num1 := 20.0
	num2 := 5.0
	resultadoEsperado := 4.0
	errorEsperado := "el divisor no puede ser 0"

	resultado, err := Dividir(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("funcion Dividir() arrojo el resultado = %v, pero el resultado esperado es %v.", resultado, resultadoEsperado)
	}

	if err != nil {
		if err.Error() != errorEsperado {
			t.Errorf("funcion Dividir() arrojo el error = %v, pero esperaba el error = %v.", err, errorEsperado)
		}
	}
}

func TestDividirMultiple(t *testing.T) {
	var tests = []struct {
		a, b float64
		want float64
		err  string
	}{
		{0, 1, 0, ""},
		{1, 1, 1, ""},
		{2, -2, -1, ""},
		{0, -1, 0, ""},
		{1, 0, 0, "el divisor no puede ser 0"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := Dividir(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("respondio %f, esperaba %f", ans, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err {
					t.Errorf("got %v, want %v", err, tt.err)
				}
			}
		})
	}
}

func TestDividirMultiple2(t *testing.T) {
	var tests = []struct {
		a, b float64
		want float64
		err  string
	}{
		{0, 1, 0, ""},
		{1, 1, 1, ""},
		{2, -2, -1, ""},
		{0, -1, 0, ""},
		{1, 0, 0, "el divisor no puede ser 0"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := Dividir(tt.a, tt.b)
			assert.Equal(t, tt.want, ans, "deben ser iguales")
			if tt.b != 0 {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
