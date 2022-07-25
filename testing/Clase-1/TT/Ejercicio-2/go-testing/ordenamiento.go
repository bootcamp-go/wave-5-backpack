package ordenamiento

func Ordenar(numeros []int) []int {

	for i := 0; i < len(numeros)-1; i++ {
		index := i
		value := numeros[i]
		for j := i + 1; j < len(numeros); j++ {
			if numeros[j] < numeros[index] {
				index = j
			}
		}
		numeros[i] = numeros[index]
		numeros[index] = value
	}

	return numeros
}
