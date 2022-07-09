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

func main()  {
	minValue, err := calculo(minimum)
	if err != nil{
		fmt.Println(err)
		return
	}
	averageValue, err := calculo(average)
	if err != nil{
		fmt.Println(err)
		return
	}
	maxValue, err := calculo(maximum)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("Promedio %.2f\n", averageValue(8, 7, 2, 1))
	fmt.Printf("Minimo %.2f\n", minValue(1, 5, 0, 4))
	fmt.Printf("Maximo %.2f\n", maxValue(7, 3, 10, 6))
}

func calculo(calculo string) (func(...float64) float64, error) {
	switch calculo {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFun, nil
	case average:
		return avarageFunc, nil
	default:
		msgError := fmt.Sprintf("La función (%s) no ha sido definida.", calculo)
		return nil, errors.New(msgError)
	}
}

func minFunc(notas ...float64) float64 {
	var min float64
	for i, nota := range notas {
		if i == 0 {
			min = nota
		}
		if nota < min {
			min = nota
		}
	}
	return min
}

func avarageFunc(notas ...float64) float64{
	var promedio float64
	var totalNotas float64

	for _, valor := range notas {
		totalNotas += valor
	}

	promedio = totalNotas / float64(len(notas))
	return promedio
}

func maxFun(notas ...float64) float64{
	var max float64
	for i, nota := range notas {
		if i == 0 {
			max = nota
		}
		if nota > max {
			max = nota
		}
	}
	return max
}
