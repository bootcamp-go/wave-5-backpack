package main

import "fmt"

func main() {
	fmt.Printf("El total de horas es: %v \n", calcSalary(120))
}

func calcSalary(minutes float64) float64 {
	var hours float64 = minutes / 60
	return hours

}

//Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
