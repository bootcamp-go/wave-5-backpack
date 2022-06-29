package main

import "fmt"

// Ejercicio 3 - Calcular salario

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad
// de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
// por mes y la categoría, y que devuelva su salario.

const (
	A string = "A"
	B string = "B"
	C string = "C"
)

func calcularSalarioEmpleado(minutos float64, categoria string) float64 {
	horasTrabajadas := minutos / 60

	switch categoria {
	case A:
		salario := 3000 * horasTrabajadas
		return salario + salario/2
	case B:
		salario := 1500 * horasTrabajadas
		return salario + salario/5
	case C:
		return 1000 * horasTrabajadas
	}

	return 0
}

func main() {
	fmt.Println("Ejercicio 3 - Calcular salario")
	fmt.Println("")

	salarioEmpleado1 := calcularSalarioEmpleado(2050, A)
	fmt.Printf("Salario del empleado 1: %.2f\n", salarioEmpleado1)

	salarioEmpleado2 := calcularSalarioEmpleado(3535, B)
	fmt.Printf("Salario del empleado 2: %.2f\n", salarioEmpleado2)

	salarioEmpleado3 := calcularSalarioEmpleado(3100, C)
	fmt.Printf("Salario del empleado 3: %.2f\n", salarioEmpleado3)
}
