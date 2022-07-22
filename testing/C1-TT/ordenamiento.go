package c1tt

func Ordenar(numeros []int) []int {
	for i := len(numeros) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if numeros[j] > numeros[j+1] {
				numeroActual := numeros[j]
				numeros[j] = numeros[j+1]
				numeros[j+1] = numeroActual
			}
		}
	}
	return numeros
}
