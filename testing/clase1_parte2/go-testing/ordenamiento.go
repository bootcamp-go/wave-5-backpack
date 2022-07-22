package claseTesting

func Ordernar(array []int) []int {
	var arrayOrdenado []int

	for i := 0; i < len(array); i++ {
		arrayOrdenado = insertarOrdenado(arrayOrdenado, array[i])
	}

	return arrayOrdenado
}

func insertarOrdenado(array []int, valor int) []int {
	if len(array) == 0 {
		return append(array, valor)
	}

	for i := 0; i < len(array); i++ {
		actual := array[i]
		if valor <= actual {
			var result = []int{valor}
			return append(result, array...)
		}
		if i == len(array)-1 {
			return append(array, valor)
		}
		siguiente := array[i+1]

		if valor >= actual && valor <= siguiente {
			return insertar(array, i+1, valor)
		}
	}

	return array
}

func insertar(array []int, index int, valor int) []int {
	cabeza := array[:index]
	cola := array[index:]

	var resultado []int
	resultado = append(resultado, cabeza...)
	resultado = append(resultado, valor)
	resultado = append(resultado, cola...)

	return resultado
}
