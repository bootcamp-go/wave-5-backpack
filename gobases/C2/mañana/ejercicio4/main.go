package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso,
requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra
función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el
cálculo que se indicó en la función anterior.
*/

const (
	minimun = "minimun"
	maximun = "maximun"
	average = "average"
)

func opMin(valores ...float64) float64 {
	var min float64
	for i, valor := range valores {
		if i == 0 {
			min = valor
		}

		if valor < min {
			min = valor
		}
	}

	return min
}

func opMax(valores ...float64) float64 {
	var max float64
	for i, valor := range valores {
		if i == 0 {
			max = valor
		}

		if valor > max {
			max = valor
		}
	}

	return max
}

func opAverage(valores ...float64) float64 {
	var promedio float64
	var totalNotas float64

	for _, valor := range valores {
		totalNotas += valor
	}

	promedio = totalNotas / float64(len(valores))
	return promedio
}

func operation(function string) (func(...float64) float64, error) {
	switch function {
	case minimun:
		return opMin, nil
	case maximun:
		return opMax, nil
	case average:
		return opAverage, nil
	default:
		msgError := fmt.Sprintf("La función con nombre (%s) no ha sido definida.", function)
		return nil, errors.New(msgError)
	}
}

func main() {
	minValue, err := operation(minimun)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageValue, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxValue, err := operation(maximun)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("✅ Promedio %.2f\n", averageValue(10, 7, 10, 10)) // 9.25
	fmt.Printf("✅ Minimo %.2f\n", minValue(1, 10, 10, 4))        // 1.00
	fmt.Printf("✅ Maximo %.2f\n", maxValue(10, 3, 10, 1))        // 10.00
}
