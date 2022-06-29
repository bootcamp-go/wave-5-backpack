package main

import "fmt"

// Ejercicio 1 - Impuestos de salario

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
// el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto
// de un salario.

// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo
// y si gana más de $150.000 se le descontará además un 10%.

func calculoImpuesto(impuesto float64) float64 {
	if impuesto > 50000 && impuesto <= 150000 {
		return (impuesto * 17) / 100
	} else if impuesto > 150000 {
		return (impuesto * 27) / 100
	} else {
		return 0
	}
}

func main() {
	fmt.Println("Ejercicio 1 - Impuestos de salario")
	fmt.Println("")

	impuestoEmpleado1 := calculoImpuesto(175150)
	fmt.Println("El impuesto del salario 1 es: ", impuestoEmpleado1)

	impuestoEmpleado2 := calculoImpuesto(65150)
	fmt.Println("El impuesto del salario 2 es: ", impuestoEmpleado2)
}
