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

func main() {

	minFunc, err := operation(MINIMUM)
	avgFunc, err := operation(AVERAGE)
	maxFunc, err := operation(MAXIMUM)
	//_, err = operation("operacionInvalida")

	var minimo = minFunc(5, 2, 3)
	var promedio = avgFunc(5, 2, 3)
	var maximo = maxFunc(5, 2, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Resultado de la operación MINIMUM: ", minimo)
	fmt.Println("Resultado de la operación AVERAGE: ", promedio)
	fmt.Println("Resultado de la operación MAXIMUM: ", maximo)
}

func operation(operacion string) (func(...float32) float32, error) {
	switch operacion {
	case MINIMUM:
		return minimo, nil
	case AVERAGE:
		return promedio, nil
	case MAXIMUM:
		return maximo, nil
	}

	return nil, errors.New("Operación no válida")

}

func minimo(valores ...float32) float32 {
	var minimo float32 = valores[0]

	for _, valor := range valores {
		if valor < minimo {
			minimo = valor
		}
	}
	return minimo
}

func maximo(valores ...float32) float32 {
	var maximo float32 = valores[0]

	for _, valor := range valores {
		if valor > maximo {
			maximo = valor
		}
	}
	return maximo
}

func promedio(valores ...float32) float32 {
	var cantidad int = len(valores)
	var acum float32 = 0
	for _, valor := range valores {
		acum += valor
	}
	return acum / float32(cantidad)
}
