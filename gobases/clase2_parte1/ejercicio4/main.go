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

func calcularMinimo(calificaciones ...float64) float64 {
	minimo := 0.0
	for _, nota := range calificaciones {
		if minimo == 0.0 {
			minimo = nota
		}
		if nota < minimo {
			minimo = nota
		}
	}
	return minimo
}

func calcularMaximo(calificaciones ...float64) float64 {
	maximo := 0.0
	for _, nota := range calificaciones {
		if maximo == 0.0 {
			maximo = nota
		}
		if nota > maximo {
			maximo = nota
		}
	}
	return maximo
}

func calcularPromedio(calificaciones ...float64) float64 {
	promedio := 0.0
	for _, nota := range calificaciones {
		promedio += nota
	}
	return promedio / float64(len(calificaciones))
}

func orquestarCalculos(calculo string) (func(valores ...float64) float64, error) {
	switch calculo {
	case minimum:
		return calcularMinimo, nil
	case maximum:
		return calcularMaximo, nil
	case average:
		return calcularPromedio, nil
	default:
		return nil, errors.New("Calculo no valido")
	}
}

func main() {
	myFunc, err := orquestarCalculos(average)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(myFunc(1, 3, 5, 7, 9))
	}
}
