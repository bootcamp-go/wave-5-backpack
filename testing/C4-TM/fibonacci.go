package fibonacci

import "errors"

func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n debe ser >= 0")
	}
	if n == 0 {
		return 0, nil
	}
	ant := 0
	res := 1
	for i := 1; i < n; i++ {
		aux := ant
		ant = res
		res += aux
	}
	return res, nil
}
