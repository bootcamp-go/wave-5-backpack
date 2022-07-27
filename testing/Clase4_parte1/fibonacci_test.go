package Clase4_parte1

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{{1, 4, 5}}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := fibonnaci(test.arg1, test.arg2)
			if got != test.want {
				t.Errorf("Test[%d]: fibonacci(%d%d) returned %d, want %d",
					i, test.arg1, test.arg2, got, test.want)
			}
		})
	}
}
