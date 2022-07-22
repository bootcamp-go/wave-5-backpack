package calculadora

// Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente

func OrdernarEnteros(ListaDesordenada []int) []int {

	//	slice :=

	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	return ListaDesordenada

}
