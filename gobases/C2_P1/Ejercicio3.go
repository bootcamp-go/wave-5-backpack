package main

import "fmt"

//Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func salaryCalculator(categoria string, minutes float64) float64 {
	switch categoria {
	case "A":
		hours := minutes / 60
		salary := hours * 3000
		fullSalary := salary + (salary * 0.50)
		return fullSalary
	case "B":
		hours := minutes / 60
		salary := hours * 1500
		fullSalary := salary + (salary * 0.20)
		return fullSalary
	case "C":
		hours := minutes / 60
		salary := hours * 1000
		return salary
	default:
		return 0
	}
}

func main() {
	salaryDetermination := salaryCalculator("C", 10678)
	fmt.Println(salaryDetermination)
}
