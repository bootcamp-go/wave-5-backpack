package fibonacci

func Fibonacci(n int) int {

	if n < 2 {
		return n
	}

	var a, b int
	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}
