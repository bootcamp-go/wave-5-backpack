package clase4testingprt1tdd

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 1}, {5, 120}}

	for i, d := range tests {
		got := FactorialMemoization(uint64(d.arg))
		if got != uint64(d.want) {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}

}
