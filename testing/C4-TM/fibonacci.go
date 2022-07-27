package fibonacci

func fibonacci(n int) int {
	ant := 0
	res := 1
	for i := 1; i < n; i++ {
		aux := ant
		ant = res
		res += aux
	}
	return res
}
