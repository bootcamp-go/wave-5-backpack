package main

import (
	"errors"
	"fmt"
)

const (
	minimun = "minimun"
	average = "average"
	maximum = "maximum"
)

func operation(tipo string, numeros ...int) (float64, error) {
	switch tipo {
	case minimun:
		return minFunc(numeros), nil
	case average:
		return avFunc(numeros), nil
	case maximum:
		return maxFunc(numeros), nil
	default:
		return 0, errors.New("Esta operacion no se encuentra")
	}

}

func minFunc(numeros []int) float64 {
	var min float64 = float64(numeros[0])

	for _, num := range numeros {
		if float64(num) < min {
			min = float64(num)
		}
	}
	return min
}

func avFunc(numeros []int) float64 {
	var total int
	for _, num := range numeros {
		total += num
	}
	return float64(total / len(numeros))
}

func maxFunc(numeros []int) float64 {
	var max float64 = float64(numeros[0])

	for _, num := range numeros {
		if float64(num) > max {
			max = float64(num)
		}
	}
	return max
}

func main() {
	fmt.Println(operation(maximum, 1, 10, 23, 8, 7))
	fmt.Println(operation(minimun, 1, 2, 10, 23, 8, 7))
	fmt.Println(operation(average, 2, 4, 6, 8, 10, 12))
	fmt.Println(operation("OTRACOSA", 1, 2, 4, 5, 6, 7))
}
