package fibo

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1

	for i := n; i > 0; i-- {
		a, b = b, b+a
	}
	return a
}
