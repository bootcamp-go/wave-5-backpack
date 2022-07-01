package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMinimum(calificaciones ...float64) (min float64) {
	for i, calificacion := range calificaciones {
		if i == 0 || calificacion < min {
			min = calificacion
		}
	}
	return
}

func opAverage(calificaciones ...float64) (promedio float64) {
	suma := .0
	for _, calificacion := range calificaciones {
		suma += calificacion
	}
	promedio = suma / float64(len(calificaciones))
	return
}

func opMaximum(calificaciones ...float64) (max float64) {
	for i, calificacion := range calificaciones {
		if i == 0 || calificacion > max {
			max = calificacion
		}
	}
	return
}

func operation(operador string) (func(calificaciones ...float64) float64, error) {
	switch operador {
	case minimum:
		return opMinimum, nil
	case average:
		return opAverage, nil
	case maximum:
		return opMaximum, nil
	default:
		return nil, errors.New("Operador no valido")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Calificación mas baja: %f\n", minValue)

	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Calificación promedio: %f\n", averageValue)

	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Calificación mas alta: %f\n", maxValue)

}
