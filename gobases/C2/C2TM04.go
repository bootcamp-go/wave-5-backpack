package main

import (
	"fmt"
)

func maximum(valores ...float64) float64 {

	numeroMayor := 0.0
	for key, valor := range valores {
		if key == 0 {
			numeroMayor = valor
		}
		if valor > numeroMayor {
			numeroMayor = valor
		}
	}
	return numeroMayor
}

func average(valores ...float64) float64 {
	suma := 0.0

	for _, valor := range valores {

		if valor >= 0 {
			suma = suma + valor
		} else {
			break
		}
	}
	return suma / float64(len(valores))
}

func minimum(valores ...float64) float64 {
	numeroMenor := 0.0
	for key, valor := range valores {
		if key == 0 {
			numeroMenor = valor
		}
		if valor < numeroMenor {
			numeroMenor = valor
		}
	}
	return numeroMenor
}

func errorOperador(valores ...float64) float64 {
	return 0.0
}

func operacionAritmetica(operador string) func(valores ...float64) float64 {
	switch operador {
	case "minimum":
		return minimum
	case "average":
		return average
	case "maximum":
		return maximum
	default:
		return errorOperador
	}
	return nil
}
func main() {

	minimo := operacionAritmetica("minimum")
	promedio := operacionAritmetica("average")
	maximo := operacionAritmetica("maximum")

	minValue := minimo(2.0, 3.0, 4.0, 1.0)
	averageValue := promedio(2.0, 3.0, 4.0, 1.0)
	maxValue := maximo(2.0, 3.0, 4.0, 1.0)

	fmt.Println("Minimo: ", minValue)
	fmt.Println("Promedio: ", averageValue)
	fmt.Println("Maximo: ", maxValue)

}
