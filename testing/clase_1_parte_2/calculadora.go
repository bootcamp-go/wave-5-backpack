package clase_1_parte_2

func Restar(numbers ...int) int {
	restar := 0
	for i, number := range numbers {
		if i == 0 {
			restar = number
			continue
		}
		restar -= number
	}
	return restar
}
