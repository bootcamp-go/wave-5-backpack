package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	resultado, error := operation(minimum)
	if error != nil {
		fmt.Println("Se ingreso un operador inexistente")
	} else {
		finalR := resultado(3, 2, 1)
		fmt.Println(finalR)
	}
}

func minFunc(valores ...int) int {
	var resultado int = math.MaxInt
	for _, valor := range valores {
		if valor < resultado {
			resultado = valor
		}
	}
	return resultado
}

func averageFunc(valores ...int) int {
	var resultado int
	var i int
	for _, valor := range valores {
		resultado += valor
		i++
	}
	return resultado / i
}

func maxFunc(valores ...int) int {
	var resultado int
	for _, valor := range valores {
		if resultado < valor {
			resultado = valor
		}
	}
	return resultado
}

func operation(operacion string) (func(valores ...int) int, error) {
	switch operacion {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, errors.New("Operacion no implementada")
	}
}
