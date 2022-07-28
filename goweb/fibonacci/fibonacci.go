package fibonacci

func fibonacci(iteraciones int) []int {
	fibo := []int{0, 1}
	var sum int

	for i := 0; i < (iteraciones - 2); i++ {
		sum = fibo[i] + fibo[i+1]
		fibo = append(fibo, sum)
	}

	return fibo
}
