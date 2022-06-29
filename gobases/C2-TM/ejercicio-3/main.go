package main

import (
	"errors"
	"fmt"
)

func main() {
	salario, err := calcularSalario(120, "C")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario es de %.2f\n", salario)
	}
}

func calcularSalario(minutosTrabajados uint, categoria string) (float64, error) {
	var salario float64
	horasTrabajadas := float64(minutosTrabajados / 60)

	switch categoria {
	case "C":
		salario = 1000 * horasTrabajadas
		return salario, nil
	case "B":
		salario = 1500 * horasTrabajadas
		salario *= 1.20
		return salario, nil
	case "A":
		salario = 3000 * horasTrabajadas
		salario *= 1.5
		return salario, nil
	default:
		return 0, errors.New("No existe la categoría")
	}
}
