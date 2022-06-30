package main

import (
	"fmt"
	"math"
)

//Ejercicio 4 - Calcular estadisticas

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calEst(operador string, valores ...float64) float64 {
	switch operador {
	case minimum:
		var res float64 = 0.0
		for i, value := range valores {
			if i == 0 {
				res = value
			} else {
				res = math.Min(res, value)
			}
		}
		return res

	case average:
		var coun float64 = 0.0
		var in int = 0
		for i, value := range valores {
			coun += value
			in = i + 1
		}
		return float64(coun / float64(in))
	case maximum:
		var res float64 = 0.0
		for i, value := range valores {
			if i == 0 {
				res = value
			} else {
				res = math.Max(res, value)
			}
		}
		return res
	}
	return 0
}

func main() {
	fmt.Println("-El minimo de la cadena corresponde a", calEst(minimum, 1, 2, 3, 48, 5))

	fmt.Println("-El maximo de la cadena corresponde a", calEst(maximum, 1, 2, 3, 48, 5))

	fmt.Println("-El promedio de la cadena corresponde a", calEst(average, 1, 2, 3, 48, 5))
}
