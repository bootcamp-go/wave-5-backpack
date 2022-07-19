package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func main() {
	clase := "B"
	clase = strings.ToUpper(clase)
	salario, err := calcularSalario(180, clase)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Para la clase: %v  el salario total del empleado es: %d\n", clase, salario)
	}
}

func calcularSalario(mins int, categoria string) (int, error) {
	if mins < 0 {
		return 0, errors.New("La cantidad de horas no puede ser negativa")
	}
	horas := float64(mins) * 0.0166667
	salario := 0.0

	switch categoria {
	case categoriaC:
		salario = 1000 * horas
	case categoriaB:
		salario = (1500 * horas)
		salario += (salario * 0.2)
	case categoriaA:
		salario = (30000 * horas)
		salario += (salario * 0.5)
	default:
		return 0, errors.New("Tipo de clase no reconocida, por favor ingresa una clase entre A, B y C")
	}
	return int(salario), nil
}
