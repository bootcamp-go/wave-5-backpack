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
	if err != nil {
		fmt.Println(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Resultado de minFunc: %.0f\n", minValue)
	fmt.Printf("Resultado de averageFunc: %.2f\n", averageValue)
	fmt.Printf("Resultado de maxFunc: %.0f\n", maxValue)
}

func operation(o string) (func(values ...int) float64, error) {
	switch o {
	default:
		return nil, errors.New("No es una operacion valida")
	case "minimum":
		return func(values ...int) (min float64) {
			min = float64(values[0])
			for _, v := range values {
				if min > float64(v) {
					min = float64(v)
				}
			}
			return min
		}, nil

	case "average":
		return func(values ...int) (ave float64) {
			var total int
			for _, v := range values {
				total += v
			}
			ave = float64(total) / float64(len(values))
			return ave
		}, nil

	case "maximum":
		return func(values ...int) (max float64) {
			max = float64(values[0])
			for _, v := range values {
				if max < float64(v) {
					max = float64(v)
				}
			}
			return max
		}, nil
	}
}
