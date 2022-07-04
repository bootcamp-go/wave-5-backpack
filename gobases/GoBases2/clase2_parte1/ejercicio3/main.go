package main

import (
	"errors"
	"fmt"
)

func calcularSalario(minutes int, category string) (float64, error) {
	if minutes < 0 {
		return 0, errors.New("Los minutos no pueden ser negativos")
	}
	switch category {
	case "A":
		return ((float64(minutes) / 60) * 3000) + (((float64(minutes) / 60) * 3000) * 0.5), nil
	case "B":
		return ((float64(minutes) / 60) * 1500) + (((float64(minutes) / 60) * 3000) * 0.2), nil
	case "C":
		return ((float64(minutes) / 60) * 1000), nil
	default:
		return 0, errors.New("No pertenece a una categoria")
	}
}

func main() {
	fmt.Println(calcularSalario(10000, "A"))
	fmt.Println(calcularSalario(10000, "B"))
	fmt.Println(calcularSalario(10000, "C"))
	fmt.Println(calcularSalario(10000, "D"))
	fmt.Println(calcularSalario(-10000, "A"))
}
