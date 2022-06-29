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

func minFunc(notas ...int) (int,error) {
	min := notas[0]

	for _, value := range notas {
		if value < min {
			min = value
		}
	}
	return min, errors.New("Hubo un error con la operacion")
}

func maxFunc(notas ...int) (int,error) {
	max := notas[0]

	for _, value := range notas {
		if value > max {
			max = value
		}
	}
	return max, errors.New("Hubo un error con la operacion")
}

func averageFunc(notas ...int) (int,error) {
	suma := 0
	for _, value := range notas {
		suma += value
	}

	return suma / len(notas), errors.New("Hubo un error con la operacion")
}

func operation(operation string) (func(enteros ...int) (int,error), error) {
	switch operation {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc,nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("Error")
	}
}

func main() {

	minValue,err := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue,err := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, err := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Calificación mas baja: %d\n", minValue)
	fmt.Printf("Calificación promedio: %d\n", averageValue)
	fmt.Printf("Calificación mas alta: %d\n", maxValue)
}
