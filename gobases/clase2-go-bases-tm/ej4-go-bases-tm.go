/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #4:  Calcular estadísticas
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		It is requested to generate a function that indicates the type of
		calculation to be performed (minimum, maximum or average) and that
		returns another function (and a message in case the calculation
		is not defined) that can be passed a number N of integers and returns
		the calculation indicated in the previous function.

		Example:
			const  (
				minimum  =  "minimum"
				average  =  "average"
				maximum  =  "maximum"
			)
			...
			minFunc  ,  err  :=  operation  (minimum)
			averageFunc, err  :=  operation  (average)
			maxFunc, err  :=  operation  (maximum)
			...
			minValue  :=  minFunc(2, 3, 3, 4, 10, 2, 4, 5)
			averageValue  :=  averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
			maxValue  :=  maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"errors"
	"fmt"
)

//	CONSTANTS
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

//	FUNCTIONS : minOp, averageOp, maxOp & operation
func minOp(valores ...float64) float64 {
	min := valores[0]
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func averageOp(valores ...float64) float64 {
	var resultado_suma float64
	for _, valor := range valores {
		resultado_suma += valor
	}
	return resultado_suma / (float64(len(valores)))
}

func maxOp(valores ...float64) float64 {
	max := valores[0]
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

func operation(operacionElegida string) (func(valores ...float64) float64, error) {
	switch operacionElegida {
	case minimum:
		return minOp, nil
	case average:
		return averageOp, nil
	case maximum:
		return maxOp, nil
	}

	return nil, errors.New("** Esta operacion no se encuentra en la lista **")
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Calculo Estadistico ||")

	// Minimo
	fmt.Println("> Minimo")
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("  Valor minimo: ", minValue)
	}

	// Maximo
	fmt.Println("> Maximo")
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("  Valor maximo: ", maxValue)
	}

	// Promedio
	fmt.Println("> Promedio")
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	} else {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("  Valor promedio: ", averageValue)
	}
}
