package main

import (
	"fmt"
	"strings"
)

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func calculoSalario(minutosTrabajados int, categoria string) float32 {
	var salario float32

	categoria = strings.ToLower(categoria)
	if categoria == "c" {
		salario = float32(minutosTrabajados) / 60 * 1000
	} else if categoria == "b" {
		salario = float32(minutosTrabajados) / 60 * 1500
		salario += salario * 0.20
	} else if categoria == "a" {
		salario = float32(minutosTrabajados) / 60 * 3000
		salario += salario * 0.50
	}

	return salario
}

func main() {
	fmt.Println(calculoSalario(100, "a"))
}
