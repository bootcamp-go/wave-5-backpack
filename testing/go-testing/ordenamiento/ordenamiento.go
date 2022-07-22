package ordenamiento

func Ordenar(numeros []int) []int {
	for i := 0; i < len(numeros); i++ {
		for j := i + 1; j < len(numeros); j++ {
			if numeros[i] > numeros[j] {
				numeros[i], numeros[j] = numeros[j], numeros[i]
			}
		}
	}
	return numeros
}
