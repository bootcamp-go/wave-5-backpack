package main

import (
	"errors"
	"fmt"
)

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
// de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
// y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido)
// que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	nada    = "nada"
)

func minValue(calificaciones ...int) int {
	minimo := calificaciones[0]

	for i := range calificaciones {

		if calificaciones[i] < minimo {
			minimo = calificaciones[i]
		}
	}

	return minimo
}

func averageValue(calificaciones ...int) int {
	var promedio int
	var notas int

	for _, calif := range calificaciones {
		notas += calif
	}

	promedio = notas / len(calificaciones)

	return promedio
}

func maxValue(calificaciones ...int) int {
	var maximo int

	for _, max := range calificaciones {
		if max > maximo {
			maximo = max
		}
	}

	return maximo
}

func operation(operacion string) (func(calificaciones ...int) int, error) {
	switch operacion {
	case "minimum":
		return minValue, nil
	case "average":
		return averageValue, nil
	case "maximum":
		return maxValue, nil
	default:
		return nil, errors.New("No hay funcion para eso")

	}

}

func main() {
	minFunc, err := operation(minimum)

	if err != nil {
		fmt.Println(err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println(minValue)
	}

	averageFunc, err := operation(average)

	if err != nil {
		fmt.Println(err)
	} else {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(averageValue)
	}

	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(maxValue)
	}

	otrFunc, err := operation(nada)

	if err != nil {
		fmt.Println(err)
	} else {
		otrFunc := otrFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(otrFunc)
	}

}
