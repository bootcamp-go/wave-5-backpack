package ejercicio2

func Ordenar(sliceEnteros []int) []int {
	var auxiliar int
	for i := 1; i < len(sliceEnteros); i++ {
		auxiliar = sliceEnteros[i]
		for j := i - 1; j >= 0 && sliceEnteros[j] > auxiliar; j-- {
			sliceEnteros[j+1] = sliceEnteros[j]
			sliceEnteros[j] = auxiliar
		}
	}
	return sliceEnteros
}
