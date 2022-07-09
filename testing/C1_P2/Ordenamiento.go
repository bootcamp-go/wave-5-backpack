package C1_P2

func Ordenamiento(data []int) []int {
	pivot := data[0]
	var order []int
	for _, value := range data {
		if value < pivot {
			pivot = value
		}
		order = append(order, pivot)
	}
	return order
}
