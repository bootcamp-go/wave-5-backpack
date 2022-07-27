package fibo

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1

	for i := n - 1; i > 0; i-- {
		a, b = b, b+a
	}
	return b
}
