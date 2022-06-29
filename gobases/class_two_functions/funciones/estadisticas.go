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
	fmt.Println("error", err)
	if err == nil {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("El mínimo valor es: ", minValue)
		fmt.Println("El máximo valor es: ", maxValue)
		fmt.Println("El promedio valor es: ", averageValue)
	}

}

func operation(typeOp string) (func(nums ...float32) float32, error) {
	switch typeOp {
	case minimum:
		return minFun, nil

	case maximum:
		return maxFunc, nil

	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("No existe la operacion especificada")
	}

}

func minFun(nums ...float32) float32 {
	var minValue float32 = 999999
	for _, num := range nums {
		if num < minValue {
			minValue = num
		}
	}

	return minValue

}

func maxFunc(nums ...float32) float32 {
	var maxValue float32 = -9999
	for _, num := range nums {
		fmt.Println("aca?", num, maxValue, num > maxValue)
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

func averageFunc(nums ...float32) float32 {
	var avgValue float32 = 0
	var plus float32 = 0
	for _, num := range nums {
		plus += num
	}
	avgValue = plus / float32(len(nums))
	return avgValue
}
