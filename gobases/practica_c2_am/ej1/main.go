package main

import "fmt"

func main() {
	salario := 151000.0
	fmt.Println("El salario es: ", salario)
	descuento, valorDescontado := impuestoSalario(salario)
	fmt.Println("Se hará un descuento del ", descuento, " del salario, que equivale a: ", valorDescontado)

	salario = 50000.0
	fmt.Println("El salario es: ", salario)
	descuento, valorDescontado = impuestoSalario(salario)
	fmt.Println("Se hará un descuento del ", descuento, " del salario, que equivale a: ", valorDescontado)
}

func impuestoSalario(salario float64) (float64, float64) {
	descuento := 0.0
	if salario > 50000.0 {
		descuento += 0.17
	}
	if salario > 150000.0 {
		descuento += 0.1
	}
	return descuento, salario * descuento
}
