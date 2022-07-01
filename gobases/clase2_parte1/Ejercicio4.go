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

func main() {

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("El minimo es : %2.f \n", minValue)
		fmt.Printf("El promedio es : %2.f \n", averageValue)
		fmt.Printf("El máximo es : %2.f \n", maxValue)
	}
}

func opMin(values ...float64) float64 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func opMax(values ...float64) float64 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func opAvg(values ...float64) float64 {
	var average float64
	for _, value := range values {
		average += value
	}
	return average / float64(len(values))
}

func operation(option string) (func(values ...float64) float64, error) {
	switch option {
	case minimum:
		return opMin, nil
	case average:
		return opAvg, nil
	case maximum:
		return opMax, nil
	default:
		return nil, errors.New("No existe la operacion")
	}
}
