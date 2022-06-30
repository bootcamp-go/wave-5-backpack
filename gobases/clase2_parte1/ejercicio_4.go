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

func minimo(calificaciones ...float32) float32 {
	min := calificaciones[0]
	for _, valor := range calificaciones {
		if valor < min {
			min = valor
		}
	}
	return min
}

func maximo(calificaciones ...float32) float32 {
	max := calificaciones[0]
	for _, valor := range calificaciones {
		if valor > max {
			max = valor
		}
	}
	return max
}

func promedio(calificaciones ...float32) float32 {
	var suma float32
	cantidad := len(calificaciones)
	for _, valor := range calificaciones {
		suma += valor
	}
	avg := suma / float32(cantidad)
	return avg
}

func tipoCalculo(calculo string) (func(calificaciones ...float32) float32, error) {
	switch calculo {
	case minimum:
		return minimo, nil
	case maximum:
		return maximo, nil
	case average:
		return promedio, nil
	}
	return nil, errors.New("El cálculo no está definido.")
}

func main() {
	res, err := tipoCalculo(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El minimo es ", res(3, 2, 4, 1, 7))
	}

	res2, err2 := tipoCalculo(maximum)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("El maximo es ", res2(3, 2, 4, 1, 7))
	}

	res3, err3 := tipoCalculo(average)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("El promedio es ", res3(3, 2, 4, 1, 7))
	}

	res4, err4 := tipoCalculo("media")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("La media es ", res4(3, 2, 4, 1, 7))
	}
}
