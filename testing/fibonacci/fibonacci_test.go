package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	lote := []uint{1, 2, 5, 3, 4}

	for _, value := range lote {
		for n := 0; n < len(lote); n++ {
			got := Fibonacci(value)
			fmt.Println(got)
		}

	}

	/* tests := []struct {
		arg  int
		want int
	}{{1, 1}, {5, 8}}

	for _, v := range tests {
		got := Fibonacci(uint(v.arg))
		if int(got) != v.want {
			t.Errorf("test fail. returned: %d expect: %d", v.arg, v.want)
		}
	} */
}
