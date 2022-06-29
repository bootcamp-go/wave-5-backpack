package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	other   = "other"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El minimo es de %f\n", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio es de %f\n", averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El máximo es de %f\n", maxFunc(2, 3, 3, 4, 1, 2, 4, 5))
	}
	otherFunc, err := operation(other)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El máximo es de %f\n", otherFunc(1, 2, 3, 4, 5))
	}
}

func operation(operation string) (func(numbers ...int) float64, error) {
	switch operation {
	case minimum:
		return minimum_func, nil
	case average:
		return average_func, nil
	case maximum:
		return maximum_func, nil

	}
	return nil, errors.New("El caso " + operation + " no está definido")
}

func minimum_func(numbers ...int) float64 {
	min_val := numbers[0]
	for _, number := range numbers {
		if min_val > number {
			min_val = number
		}
	}
	return float64(min_val)
}

func average_func(numbers ...int) float64 {
	var average_val float64 = 0
	for _, number := range numbers {
		average_val += float64(number)
	}
	average_val = average_val / float64(len(numbers))
	return average_val
}

func maximum_func(numbers ...int) float64 {
	max_val := numbers[0]
	for _, number := range numbers {
		if max_val < number {
			max_val = number
		}
	}
	return float64(max_val)
}
