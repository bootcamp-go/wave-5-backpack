package fibonacci

func fibonacci(number int) int {
	if number <= 1 {
		return number
	} else {
		return fibonacci(number-2) + fibonacci(number-1)
	}
}
