package Clase12

func Ordenamiento(values []int) []int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[i] {
				aux := values[j]
				values[j] = values[i]
				values[i] = aux
			}
		}
	}
	return values
}
