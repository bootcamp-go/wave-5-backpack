package ordenamiento

func Ordenamiento(valores []int) []int {
	var valMenor, posMenor, aux int
	for i := 0; i < len(valores); i++ {
		valMenor = valores[i]
		posMenor = i

		for j := i + 1; j < len(valores); j++ {
			if valores[j] < valMenor {
				valMenor = valores[j]
				posMenor = j
			}
		}

		if i != posMenor {
			aux = valores[i]
			valores[i] = valMenor
			valores[posMenor] = aux
		}
	}

	return valores
}
