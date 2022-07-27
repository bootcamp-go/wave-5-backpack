package fibonacci

func fibonacci(n int) int {
	serie := []int{0, 1, 1}

	if n <= 2 {
		return serie[n]
	}

	for i := 3; i <= n; i++ {
		suma := serie[i-1] + serie[i-2]
		serie = append(serie, suma)
	}
	return serie[n]
}
