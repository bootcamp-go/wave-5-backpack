package clase4testingprt1tdd

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{1, 1}, {2, 1}, {3, 2}, {5, 5}, {8, 21}, {13, 233}}

	for i, d := range tests {
		got := FibonacciRecursion(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
