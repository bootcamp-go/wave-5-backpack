package ordenamiento

func ordenamientoBurbuja(listaInt []int) []int {
	var auxiliar int
	for i := 0; i < len(listaInt); i++ {
		for j := 0; j < len(listaInt); j++ {
			if (listaInt)[i] < (listaInt)[j] {
				auxiliar = (listaInt)[i]
				(listaInt)[i] = (listaInt)[j]
				(listaInt)[j] = auxiliar
			}
		}
	}
	return listaInt
}
