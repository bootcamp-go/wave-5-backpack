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

func opMin(values ...float64) float64 {
	var resultado float64 = values[0]
	for _, value := range values {
		if resultado > value {
			resultado = value
		}
	}

	return resultado
}

func opAve(values ...float64) float64 {
	var resultado float64
	for _, value := range values {
		resultado += value
	}

	return (resultado / float64(len(values)))
}

func opMax(values ...float64) float64 {
	var resultado float64 = values[0]
	for _, value := range values {
		if resultado < value {
			resultado = value
		}
	}

	return resultado
}

func operation(op string) (func(...float64) float64, error) {

	switch op {
	case minimum:
		return opMin, nil
	case average:
		return opAve, nil
	case maximum:
		return opMax, nil
	}

	return nil, errors.New("La operacion no esta permitida")
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("no se puedo calcular el valor minimo")
		return
	}

	averageFunc, err := operation("aver")
	if err != nil {
		fmt.Println("no se puedo calcular el valor promedio")
		return
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println("no se puedo calcular el valor maximo")
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("min: %.2f\nprom: %.2f\nmax: %.2f\n", minValue, averageValue, maxValue)
}
