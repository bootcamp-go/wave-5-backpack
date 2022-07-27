package fibonacci

import "testing"

// FUNCIONALIDAD: Calcula el n-ésimo número de Fibonacci
// CASOS DE PRUEBA:
// 1. n = 0, 1
// 2. n = 1, 1
// 3. n = 2, 1
// 4. n = 3, 2
// 5. n = 4, 3
// 6. n = 5, 5
// 7. n = 50, 12586269025

type casosTest struct {
	param int
	res   int64
}

func TestFibonacci(t *testing.T) {
	tests := []casosTest{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{50, 12586269025},
	}

	for i, test := range tests {
		res := Fibonacci(test.param)
		if res != test.res {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, expected %d", i, test.param, res, test.res)
		}
	}
}
