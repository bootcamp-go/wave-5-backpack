package c4fibonacci

import (
	"testing"
)

func TestFibonacciLastNumb(t *testing.T) {

	tests := []struct {
		n    int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}}

	for n, v := range tests {
		actual := FibonacciLastNumb(v.n)
		if actual != v.want {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d", n, v.n, actual, v.want)
		}
	}
}
