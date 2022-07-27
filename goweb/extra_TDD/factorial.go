package extratdd

/*func factorial1(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial1(number-1)
}*/

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	f := 1

	for i := 1; i <= number; i++ {
		f *= i
	}

	return f
}
