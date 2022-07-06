// Ejercicio 3 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
package main

import "fmt"

func calculoHoras(minutos, valorHora int) float64 {
	return float64(minutos / 60 * valorHora)
}
func salario(minutos int, categoria string, calculoHoras func(minutos, valorHora int) float64) float64 {
	switch categoria {
	case "A":
		return calculoHoras(minutos, 1000)
	case "B":
		return calculoHoras(minutos, 1500) * 1.20

	case "C":
		return calculoHoras(minutos, 3000) * 1.50

	default:
		return 0
	}
}

func main() {

	fmt.Println("Ingrese la categoria ")
	var cat string 

	fmt.Scanln(&cat)
	fmt.Println("Ingrese la cantidad de minutos ")
	var minutos int 

	fmt.Scanln(&minutos)
	fmt.Printf("El sueldo es: %.2f \n", salario(minutos, cat, calculoHoras))
}
