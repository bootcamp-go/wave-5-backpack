package main

import (
	"errors"
	"fmt"
)

func max(valores []float64) float64 {
	result := valores[0]
	for _, valor := range valores {
		if valor > result {
			result = valor
		}
	}

	return result
}

func min(valores []float64) float64 {
	result := valores[0]
	for _, valor := range valores {
		if valor < result {
			result = valor
		}
	}
	return result
}

func promedio(valores []float64) float64 {
	var result float64
	count := 0.0
	for _, valor := range valores {
		count++
		result += valor
	}
	result = result / count
	return result
}

func orquestador(valores []float64, operacion func(valores []float64) float64) float64 {
	return operacion(valores)
}

func calculo(tipoCalculo string, valores ...float64) (float64, error) {
	switch tipoCalculo {
	case "max":
		return orquestador(valores, max), nil
	case "min":
		return orquestador(valores, min), nil
	case "promedio":
		return orquestador(valores, promedio), nil

	}

	return 0, errors.New("No ingresaste los datos correctamente")
}

func main() {

	res, err := calculo("max", 3.0, 4.0, 5.0, 10.0)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	} else {
		fmt.Printf("El resulto de la operación es: %v \n", res)
	}

}
