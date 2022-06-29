package main

import "fmt"

func impuestoDeSalario(salario int) float64 {
	impuesto := .0
	if salario > 50000 {
		impuesto += .17
	}
	if salario > 150000 {
		impuesto += .10
	}
	return impuesto * float64(salario)
}

func main() {
	sueldo := 55000
	impuesto := impuestoDeSalario(sueldo)
	fmt.Printf("El impuesto a pagar es: $%f\n", impuesto)
}
