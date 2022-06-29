package main

import "fmt"

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func main() {
	c := calcularSalario(9600, "C")
	fmt.Printf("Salario categoria C: %v\n", c)

	b := calcularSalario(9600, "B")
	fmt.Printf("Salario categoria B: %v\n", b)

	a := calcularSalario(9600, "A")
	fmt.Printf("Salario categoria A: %v\n", a)
}

const (
	C = "C"
	B = "B"
	A = "A"
)

func calcularSalario(min int, categoria string) float64 {
	var salario float64
	horas := float64(min) / 60

	switch categoria {
	case C:
		salario = horas * 1000
	case B:
		salario = horas * 1500
		salario += salario * 0.20 // suma %20 del salario mensual
	case A:
		salario = horas * 3000
		salario += salario * 0.50 // suma %50 del salario mensual
	}

	return salario
}
