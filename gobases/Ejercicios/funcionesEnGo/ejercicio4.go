package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

func calcularEstadisticas(calculo string) {
	//Ejercicio 4
	minValue, err := estadistica(calculo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s nota: %.2f\n", calculo, minValue(1, 2, 3, 4, 5))
	}
}

func estadistica(operation string) (func(notas ...int) float32, error) {
	//Ejercicio 4
	switch operation {
	case MINIMUM:
		return minFunc, nil
	case AVERAGE:
		return aveFunc, nil
	case MAXIMUM:
		return maxFunc, nil
	default:
		return nil, errors.New("Calculo no definido")
	}
}

func minFunc(notas ...int) float32 {
	//Ejercicio 4
	var min int

	for i, nota := range notas {
		if i == 0 {
			min = nota
		} else if nota < min {
			min = nota
		}
	}
	return float32(min)
}

func aveFunc(notas ...int) float32 {
	//Ejercicio 4
	countNotas := 0
	sumNotas := 0

	for _, nota := range notas {
		sumNotas += nota
		countNotas++
	}
	return float32(sumNotas / countNotas)
}

func maxFunc(notas ...int) float32 {
	//Ejercicio 4
	var max int

	for i, nota := range notas {
		if i == 0 {
			max = nota
		} else if nota > max {
			max = nota
		}
	}
	return float32(max)
}
