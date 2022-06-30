/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

var (
	catA = "catA"
	catB = "catB"
	catC = "catC"
)

func calcularSalario(minTrabajados float64, cat string) float64 {
	hrTrabajadas := minTrabajados / 60.00
	switch cat {
	case catA:
		return hrTrabajadas*3000.00*1.5
	case catB:
		return hrTrabajadas*1500.00*1.2
	case catC:
		return hrTrabajadas*1000.00
	}
	return 0
}

func main(){
	fmt.Println(calcularSalario(60.0, "catB"))
}