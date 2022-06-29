package main

import (
	"errors"
	"fmt"
)

var (
	minutos int
)

func main() {
	minutos = 300
	s, err := trabajo(minutos, "A")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Salario: $%.2f\n", s)
	}
}

func trabajo(minutos int, categoria string) (salario float64, err error) {
	if minutos < 0 {
		return 0, errors.New("No se admiten valores negativos")
	}

	horas := float64(minutos) / 60

	switch categoria {
	default:
		return 0, errors.New("Categoria invalida")
	case "A":
		salario = horas * 3000
		salario += salario * 0.5
	case "B":
		salario = horas * 1500
		salario += salario * 0.2
	case "C":
		salario = horas * 1000
	}
	return salario, nil
}
