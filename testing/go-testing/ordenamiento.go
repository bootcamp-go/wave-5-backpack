package calculadora

// Función para ordenar un slice de enteros de forma ascendente
func Ordenamiento(nums ...int) []int {
	var validar int
	for i := 1; i < len(nums); i++ {
		// Se obtiene el segundo elemento del slice
		validar = nums[i]
		// Se intercalan los valores si el valor de la derecha es mayor que el valor de la izquierda
		for j := i - 1; j >= 0 && nums[j] > validar; j-- {
			nums[j+1] = nums[j]
			nums[j] = validar
		}
	}
	return nums
}
