package fibonacci

func Fibonacci(number int) int {
	if number == 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	f := []int{0, 1}

	for i := 2; i <= number; i++{
		j := len(f) - 1
		k := f[j] + (f[j - 1])
		f = append(f, k)
	}

	return f[number]
}