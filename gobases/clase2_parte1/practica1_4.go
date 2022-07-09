package main

import (
	"fmt"
	"errors"
)

//Ejercicio 4 - Calcular estadisticas
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(values ...float64) float64 {

	var min float64
	for i, value := range values {
		if i == 0 {
			min = value
		}
		if value < min {
			min = value
		}
	}

	return min

}

func averageFunc(values ...float64)  float64 {
	var result float64
	for _, value := range values {
		result += value
	}
	return result / float64(len(values))
}

func maxFunc(values ...float64) float64 {

	var max float64
	for i, value := range values {
		if i == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}

	return max

}

func operation(operationType string) (func(...float64) float64, error) {
	switch operationType {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, errors.New("La función no existe")
	}
}

func main() {
	
	min, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}

	avrg, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	max, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(min(1, 5, 3, 7, 2, -1))
	fmt.Println(avrg(1, 5, 3, 7, 2))
	fmt.Println(max(1, 5, 3, 7, 2))

}