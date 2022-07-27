package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}}

	for i, d := range tests {
		got := fibonacci(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, want %d", i, d.arg, got, d.want)
		}
	}
}
