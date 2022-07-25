package calculadora

import (
	"fmt"
	"sort"
)

func Restar(n1, n2 int) int {
	return n1 - n2
}

func OrdenarSlice(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func Dividir(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("el denominador no puede ser 0")
	}
	res := n1 / n2
	return res, nil
}
