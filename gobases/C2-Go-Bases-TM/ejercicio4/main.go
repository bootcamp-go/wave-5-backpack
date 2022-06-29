package main

import (
	"fmt"
)

// Ejercicio 4 - Calcular estadísticas

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
// de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus
// calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
// (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el
// cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el
// cálculo que se indicó en la función anterior

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(calificaciones ...int) float64 {
	primeraCalificacion := true
	calificacion := 0
	calif := 0

	for _, v := range calificaciones {
		// Validamos que las calificaciones sean entre 1 y 10
		if v < 0 {
			calif = 0
		} else if v > 10 {
			calif = 10
		} else {
			calif = v
		}

		if primeraCalificacion {
			// Obtenemos la primera calificación de referencia
			calificacion = calif
			primeraCalificacion = false
		} else if calif < calificacion {
			// Validamos las demás calificaciones
			calificacion = calif
		}
	}
	return float64(calificacion)
}

func averageFunc(calificaciones ...int) float64 {
	sumatoria := 0
	totalMaterias := 0

	for _, v := range calificaciones {
		if v < 0 {
			sumatoria += 0
		} else if v > 10 {
			sumatoria += 10
		} else {
			sumatoria += v
		}
		totalMaterias++
	}

	return float64(sumatoria / totalMaterias)
}

func maxFunc(calificaciones ...int) float64 {
	primeraCalificacion := true
	calificacion := 0
	calif := 0

	for _, v := range calificaciones {
		// Validamos que las calificaciones sean entre 1 y 10
		if v < 0 {
			calif = 0
		} else if v > 10 {
			calif = 10
		} else {
			calif = v
		}

		if primeraCalificacion {
			// Obtenemos la primera calificación de referencia
			calificacion = calif
			primeraCalificacion = false
		} else if calif > calificacion {
			// Validamos las demás calificaciones
			calificacion = calif
		}
	}
	return float64(calificacion)
}

func calculoCalificaciones(operacion string) func(calificaciones ...int) float64 {
	switch operacion {
	case minimum:
		return minFunc
	case average:
		return averageFunc
	case maximum:
		return maxFunc
	}

	return nil
}

func main() {
	fmt.Println("Ejercicio 4 - Calcular estadísticas")
	fmt.Println("")

	oper_min := calculoCalificaciones(minimum)
	func_min := oper_min(6, 8, 7, 9, 10)
	fmt.Println("Mínimo de calificaciones: ", func_min)

	oper_aver := calculoCalificaciones(average)
	func_aver := oper_aver(6, 8, 7, 9, 10)
	fmt.Println("Promedio de calificaciones: ", func_aver)

	oper_max := calculoCalificaciones(maximum)
	func_max := oper_max(6, 8, 7, 9, 10)
	fmt.Println("Mínimo de calificaciones: ", func_max)
}
