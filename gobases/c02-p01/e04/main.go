package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

func minFunc(notas ...float64) float64 {
	var min float64 = notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min
}

func averageFunc(notas ...float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func maxFunc(notas ...float64) float64 {
	var max float64 = notas[0]
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max
}

func operation(operador string) (func(notas ...float64) float64, error) {
	switch operador {
	case MINIMUM:
		return minFunc, nil
	case AVERAGE:
		return averageFunc, nil
	case MAXIMUM:
		return maxFunc, nil
	}
	return nil, errors.New("Ha ingresado una operación inválida")
}

func main() {

	fmt.Println("Las notas a considerar son:", 2, 3, 3, 4, 10, 2, 4, 5)

	minFunc, minErr := operation(MINIMUM)
	if minErr != nil {
		fmt.Println(minErr)
	} else {
		fmt.Println("El valor mínimo es:", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}

	averageFunc, averageErr := operation(AVERAGE)
	if averageErr != nil {
		fmt.Println(averageErr)
	} else {
		fmt.Println("El valor promedio es:", averageFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}

	maxFunc, maxErr := operation(MAXIMUM)
	if maxErr != nil {
		fmt.Println(maxErr)
	} else {
		fmt.Println("El valor promedio es:", maxFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}

	_, err := operation("")
	if err != nil {
		fmt.Println(err)
	}

}
