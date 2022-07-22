package calculadora

func MetodoInsercao(vetor []int) []int {
	i := 0
	j := 1
	aux := 0
	n := len(vetor)

	for j < n {
		aux = vetor[j]
		i = j - 1
		for (i >= 0) && (vetor[i] > aux) {
			vetor[i+1] = vetor[i]
			i--
		}
		vetor[i+1] = aux
		j++
	}
	return vetor
}
