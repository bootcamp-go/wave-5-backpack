package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		num      int
		sucesion string
		total    int
	}{
		{num: 15, sucesion: "0 1 1 2 3 5 8 13 21 34 55 89 144 233 377", total: 986},
		{num: 5, sucesion: "0 1 1 2 3", total: 7},
	}

	for _, value := range tests {
		gotS, gotT := Fibonnaci(value.num)
		if gotS != value.sucesion || gotT != value.total {
			t.Errorf("Failed, expected %d, but got %d", value.total, gotT)
		}
	}
}
