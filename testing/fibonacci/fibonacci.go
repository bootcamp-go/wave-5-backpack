package fibonacci

/*
func fibonacci(targetNumber int) int {
	if targetNumber == 0 || targetNumber == 1 {
		return 0
	}
	if targetNumber == 2 {
		return 1
	}
	return fibonacci(targetNumber-2) + fibonacci(targetNumber-1)
}
*/

func fibonacci(targetNumber int) int {
	if targetNumber == 1 {
		return 0
	}
	if targetNumber == 2 {
		return 1
	}
	pair := []int{0, 1}
	for i := 3; i <= targetNumber; i++ {
		pair = []int{pair[1], pair[1] + pair[0]}
	}
	return pair[1]
}
