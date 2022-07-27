package tdd

import "testing"

func TestFibonnaci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{6, 8},
		{10, 55},
	}

	for i, te := range tests {
		got := fibonacci(te.arg)
		if got != te.want {
			t.Errorf("Error en la iteración %d, fibonnaci(%d) returned %d inteado of %d", i+1, te.arg, got, te.want)
		}
	}
}
