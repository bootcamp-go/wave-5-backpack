package main

import (
	"fmt"
)

/*Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	1. Si es categoría C, su salario es de $1.000 por hora
	2. Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	3. Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func salario(minutos float64, cat string) (float64, string) {
	horas := minutos / 60
	switch cat {
	case "A":
		return horas * 3000 * 1.5, ""
	case "B":
		return horas * 1500 * 1.2, ""
	case "C":
		return horas * 1000, ""
	default:
		return 0, "Los datos son incorrectos"
	}
}

func main() {
	var minutosTrabajados float64 = 480 // 8 hrs * 60 minutos
	var categoria string = "B"

	salario, msg := salario(minutosTrabajados, categoria)
	if msg != "" {
		fmt.Println(msg)
	} else {
		fmt.Printf("Categoria: %s, Sueldo: %0.2f\n", categoria, salario)
	}
}
