package go_testing

import "errors"

func Restar(a, b int) int {
	return a - b
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("the denominator number should not be equal to 0")
	}
	return int(num / den), nil
}
