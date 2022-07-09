package main

import (
	"errors"
	"fmt"
)

const (
	minimun string = "minumun"
	maximun string = "maximun"
	average string = "average"
)

func statistics(category string, values ...int) (float64, error) {
	switch category {
	case minimun:
		var minimun int
		for _, result := range values {
			pivote := values[0]
			if result < pivote {
				pivote = result
				minimun = int(pivote)
			}
		}
		return float64(minimun), nil
	case maximun:
		var maximun int
		for _, result := range values {
			pivote := values[0]
			if result > pivote {
				pivote = result
				maximun = pivote
			}
		}
		return float64(maximun), nil
	case average:
		var result float64
		for _, suma := range values {
			if suma < 0 {
				return 0, errors.New("Negative values not acepted")
			}
			result += float64(suma)
		}
		return result / float64(len(values)), nil
	default:
		return 0, errors.New("This operations is not defined")
	}
}

func main() {
	operations, err := statistics(minimun, 4, 6, 8, 9, 10, 2)
	if err != nil {
		fmt.Println("Error presentado:", err)
	} else {
		fmt.Println("El resutado de la operación es:", operations)
	}
}
