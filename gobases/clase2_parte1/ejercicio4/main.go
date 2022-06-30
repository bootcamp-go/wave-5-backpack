package main

import (
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minValue, errMin := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	maxValue, errMax := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	averageValue, errAvg := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if errMin != nil {
		fmt.Println(errMin)
	} else {
		fmt.Printf("Calificacion mas baja: %d\n", minValue)
	}

	if errMax != nil {
		fmt.Println(errMax)
	} else {
		fmt.Printf("Calificacion mas alta: %d\n", maxValue)
	}

	if errAvg != nil {
		fmt.Println(errAvg)
	} else {
		fmt.Printf("Calificacion promedio: %d\n", averageValue)
	}
}

func minFunc(notas ...int) (int, error) {
	min := notas[0]
	for _, value := range notas {
		if value < min {
			min = value
		}
	}
	return min, nil
}

func maxFunc(notas ...int) (int, error) {
	max := notas[0]
	for _, value := range notas {
		if value > max {
			max = value
		}
	}
	return max, nil
}

func averageFunc(notas ...int) (int, error) {
	suma := 0
	for _, value := range notas {
		suma += value
	}
	return suma / len(notas), nil
}
