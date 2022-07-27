package clase4_parte1

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{{0, 1, 1}, {1, 1, 2}}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			got := fibonacci(test.arg1, test.arg2)
			if got != test.want {
				t.Errorf("Test [%d]: fibonacci(%d,%d) returned %d, want %d",
					idx, test.arg1, test.arg2, got, test.want)
			}
		})

	}
}
