package testing_go

func OrdenAscendente(enteros []int) []int {
	for i := 0; i < len(enteros); i++ {
		var entOrden = i
		for j := i; j < len(enteros); j++ {
			if enteros[j] < enteros[entOrden] {
				entOrden = j
			}
		}
		enteros[i], enteros[entOrden] = enteros[entOrden], enteros[i]
	}
	return enteros
}
