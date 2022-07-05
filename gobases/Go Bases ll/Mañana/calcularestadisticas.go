package main

import "fmt"

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

func minFunc(notas ...float64) float64 {
	var minimo float64 = 1000

	for _, nota := range notas {
		if nota < minimo {
			minimo = nota
		}
	}
	return minimo
}

func minAverage(notas ...float64) float64 {
	var sumaNotas float64
	var numNotas float64 = float64(len(notas))
	for _, nota := range notas {
		sumaNotas += nota
	}
	if numNotas == 0 {
		return 0
	} else {
		return sumaNotas / numNotas
	}
}

func maxFunc(notas ...float64) float64 {
	var maximo float64

	for _, nota := range notas {
		if nota > maximo {
			maximo = nota
		}
	}
	return maximo
}

func operacionAritmetica(operador string) func(notas ...float64) float64 {

	switch operador {
	case MINIMUM:
		return minFunc
	case AVERAGE:
		return minAverage
	case MAXIMUM:
		return maxFunc
	}
	fmt.Println("Opcion incorrecta")
	return nil
}

func main() {
	oper := operacionAritmetica(MINIMUM)
	resultado := oper(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(resultado)
}
