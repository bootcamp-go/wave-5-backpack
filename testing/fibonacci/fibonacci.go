package fibonacci

/* func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
} */

func Fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint
	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}
