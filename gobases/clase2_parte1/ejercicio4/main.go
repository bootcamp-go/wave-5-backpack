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

func minimo(valores ...float64) float64 {

	menorPromedio := 0.0
	for i, valor := range valores {
		if i == 0 {
			menorPromedio = valor
			i++
		} else {
			if menorPromedio > valor {
				menorPromedio = valor
			}
		}
	}
	return menorPromedio
}

func averages(valores ...float64) float64 {
	promedioTotal := 0.0
	divisor := float64(len(valores))

	for _, valor := range valores {
		promedioTotal += valor
	}

	return promedioTotal / divisor
}
func maximo(valores ...float64) float64 {
	mayorPromedio := 0.0
	for i, valor := range valores {
		if i == 0 {
			mayorPromedio = valor
			i++
		} else {
			if mayorPromedio < valor {
				mayorPromedio = valor
			}
		}
	}
	return mayorPromedio
}
func operacion(operador string) (func(valores ...float64) float64, error) {
	switch operador {
	case "minimum":
		return minimo, nil
	case "average":
		return averages, nil
	case "maximum":
		return maximo, nil
	}
	return nil, errors.New("no existe categoria")
}

func main() {

	a, err := operacion("average")

	if err == nil {
		res := a(1, 2, 3, 8)
		fmt.Println("resultado: ", res)
	} else {
		fmt.Println(err)
	}

}
