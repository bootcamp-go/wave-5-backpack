package main

import (
	"fmt"
)

const (
	Minimun = "minimun"
	Average = "average"
	Maximun = "maximun"
)

func minimun(nums ...float64) float64 {
	var min float64
	for i, value := range nums {
		if i == 0 {
			min = value
		} else {
			if value < min {
				min = value
			} else {
				continue
			}
		}
	}
	return min
}

func maximun(nums ...float64) float64 {
	var max float64
	for i, value := range nums {
		if i == 0 {
			max = value
		} else {
			if value > max {
				max = value
			} else {
				continue
			}
		}
	}
	return max
}

func average(nums ...float64) float64 {
	var avg float64
	for _, value := range nums {
		avg += value
	}
	return avg / float64(len(nums))
}

func selectFunc(operacion string) func(nums ...float64) float64 {
	switch operacion {
	case Minimun:
		return minimun
	case Maximun:
		return maximun
	case Average:
		return average
	}
	return nil
}

func main() {
	minValue := selectFunc("Hola")
	if minValue == nil {
		fmt.Printf("La funcion no existe\n")
	} else {
		min := minValue(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Printf("%v\n", min)
	}

}
