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

func minFunc(values ...float64) float64 {

	minNumber := 99999999999999.0

	for _, value := range values {
		if value < minNumber {
			minNumber = value
		}
	}
	return float64(minNumber)

}

func averageFunc(values ...float64) float64 {

	cantidadNumeros := float64(len(values))
	var sum float64
	var average float64

	for _, value := range values {
		sum += value
	}

	average = sum / cantidadNumeros
	return average
}

func maxFunc(values ...float64) float64 {
	maxNumber := 0.0

	for _, value := range values {
		if value > maxNumber {
			maxNumber = value
		}
	}
	return float64(maxNumber)

}

func funcionDeCalculo(tipoCalculo string) (func(values ...float64) float64, error) {

	switch tipoCalculo {
	case "minimum":
		return minFunc, nil
	case "average":
		return averageFunc, nil
	case "maximum":
		return maxFunc, nil
	}

	return nil, errors.New("No se ha indicado el calculo o se encuentra mal ingresado")
}

func main() {
	opera, err := funcionDeCalculo("average")
	if err != nil {
		fmt.Println(err)
	} else {
		r := opera(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Printf("El resultado de la operación es: %v \n", r)
	}

}
