package clase3tddfig

func fibonacci(n int) int {
	a := 0
	b := 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	for i := 2; i < n+1; i++ {
		newNumber := a + b
		a = b
		b = newNumber

	}
	return b
}
