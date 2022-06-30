/*
	Ejercicio 3 - Calcular salario
	
	Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad 
	de horas trabajadas por mes y la categoría.

	Si es categoría C, su salario es de $1.000 por hora
	Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados 
	por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

func main() {
	minutosTrabajadosMes := 60.0
	categoria := "A"
	fmt.Println("El cálculo del salario es: ", calculoSalario(minutosTrabajadosMes, categoria))
}

func calculoSalario(minutos float64, categoria string) float64 {
	switch categoria {
	case "A":
		return (3000.0 * calculoHoras(minutos)) * 1.5
	case "B":
		return (1500.0 * calculoHoras(minutos)) * 1.2 
	case "C":
		return 1000.0 * calculoHoras(minutos)
	}
	return 0;
}

func calculoHoras(minutos float64) float64 {
	return minutos/60.0
}