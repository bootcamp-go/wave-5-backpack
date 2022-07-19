package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "Mínimo"
	MAXIMO   = "Máximo"
	PROMEDIO = "Promedio"
)

func main() {

	calculo, err := tipoCalculo(MINIMO)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El mínimo es: ", calculo(1, 2, 3, 4, 5))
	}

	calculo, err = tipoCalculo(MAXIMO)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El máximo es: ", calculo(1, 2, 3, 4, 5))
	}

	calculo, err = tipoCalculo(PROMEDIO)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", calculo(1, 2, 3, 4, 5))
	}

	calculo, err = tipoCalculo("Otro")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", calculo(1, 2, 3, 4, 5))
	}

}

func tipoCalculo(tipo string) (func(enteros ...int) float64, error) {
	switch tipo {
	case MINIMO:
		return minimo, nil
	case MAXIMO:
		return maximo, nil
	case PROMEDIO:
		return promedio, nil
	default:
		return nil, errors.New("No existe el tipo de cálculo: " + tipo)
	}
}

func minimo(enteros ...int) float64 {
	var minimo int
	for i, entero := range enteros {
		if i == 0 {
			minimo = entero
		} else {
			if entero < minimo {
				minimo = entero
			}
		}
	}
	return float64(minimo)
}

func maximo(enteros ...int) float64 {
	var maximo int
	for i, entero := range enteros {
		if i == 0 {
			maximo = entero
		} else {
			if entero > maximo {
				maximo = entero
			}
		}
	}
	return float64(maximo)
}

func promedio(enteros ...int) float64 {
	var suma int
	for _, entero := range enteros {
		suma += entero
	}
	return float64(suma / len(enteros))
}
