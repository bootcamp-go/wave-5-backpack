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
	funcion, err := operation(maximum)

	if err != nil {
		fmt.Println(err)
	} else {
		calculo, operacion := funcion(2, 3, 3, 4, 10, 2, 4, 5)

		fmt.Printf("El valor %s es %.2f\n", operacion, calculo)
	}
}

func operation(operacion string) (func(notas ...int) (float64, string), error) {
	switch operacion {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return func(notas ...int) (float64, string) { return 0.0, "" }, errors.New("No existe la operación especificada")
	}
}

func minFunc(notas ...int) (float64, string) {
	minimo := notas[0]

	for _, nota := range notas {
		if nota < minimo {
			minimo = nota
		}
	}
	return float64(minimo), "minimo"
}

func averageFunc(notas ...int) (float64, string) {
	var promedio float64
	for _, nota := range notas {
		promedio += float64(nota)
	}
	return promedio / float64(len(notas)), "promedio"
}

func maxFunc(notas ...int) (float64, string) {
	maximo := notas[0]

	for _, nota := range notas {
		if nota > maximo {
			maximo = nota
		}
	}
	return float64(maximo), "maximo"
}
