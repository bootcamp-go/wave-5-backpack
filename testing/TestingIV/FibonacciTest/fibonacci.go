package fibonaccitest

import "errors"

func Fibonacci(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("el numero debe ser mayor a 0")
	}
	if n == 0 {
		return 0, nil
	}
	if n < 2 {
		return n, nil
	}

	fst := 0
	sd := 1
	conta := fst
	for i := 1; i < n; i++ {
		conta = fst
		fst = sd
		sd += conta
	}

	return sd, nil

	//return Fibonacci(n-1),nil + Fibonacci(n-2)
}
