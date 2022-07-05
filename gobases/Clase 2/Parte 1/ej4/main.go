package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	MAXIMO   = "maximo"
	PROMEDIO = "promedio"
)

type operacionValida func(...float64) float64

func obtenerMinimo(params ...float64) float64 {

	var minimo float64 = params[0]
	for i := 1; i < len(params); i++ {
		if minimo > params[i] {
			minimo = params[i]
		}
	}
	return minimo
}

func obtenerMaximo(params ...float64) float64 {
	var maximo float64 = params[0]
	for i := 1; i < len(params); i++ {
		if maximo < params[i] {
			maximo = params[i]
		}
	}
	return maximo
}

func sumaDatos(params ...float64) float64 {
	var sum float64
	for _, num := range params {
		sum += num
	}
	return sum
}

func promediarDatos(notas ...float64) float64 {
	return sumaDatos(notas...) / float64(len(notas))
}

func selectorOperacion(operacion string) (operacionValida, error) {
	switch operacion {
	default:
		return nil, errors.New("Operacion no valida")
	case MINIMO:
		return obtenerMinimo, nil
	case MAXIMO:
		return obtenerMaximo, nil
	case PROMEDIO:
		return promediarDatos, nil
	}
}

func calculadora(operacion string, params ...float64) (float64, error) {
	operacionSeleccionada, err := selectorOperacion(operacion)
	if err != nil {
		return 0, err
	}
	return operacionSeleccionada(params...), nil
}

func main() {
	operacion := "promedio"
	datos := []float64{5, 0, 3, 2, 4, 10}
	res, err := calculadora(operacion, datos...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
