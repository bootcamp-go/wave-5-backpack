package c4fibonacci

func FibonacciLastNumb(n int) int {

	if n < 2 {
		return n
	}

	return FibonacciLastNumb(n-1) + FibonacciLastNumb(n-2)

}
