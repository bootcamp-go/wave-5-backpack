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

func sumarElementos(values ...int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}

func contarElementos(values ...int) int {
	return len(values)

}

func opeMin(values ...int) int {
	min := int(values[0])
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}

func opeMax(values ...int) int {
	max := int(values[0])
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return max
}
func opeAve(values ...int) int {
	var acumulador int
	for _, value := range values {
		acumulador = acumulador + value
	}
	fmt.Println(acumulador)
	fmt.Println(len(values))

	return (acumulador / len(values))
}

func controlador(operacion string) (func(values ...int) int, error) {
	switch operacion {
	case minimum:
		return opeMin, nil
	case average:
		return opeAve, nil
	case maximum:
		return opeMax, nil
	}
	return nil, errors.New("EXISTE LA OPERACION INGRESADA")

}

func main() {

	r, _ := controlador(average)
	j := r(4, 4, 4, 4)
	fmt.Println("Su resultado es:", j)

}
