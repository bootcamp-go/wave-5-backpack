package main

import "fmt"

func main() {
	salario := 40000.00

	fmt.Printf("El salario despues de impuestos es de $%.2f\n", calcularImpuestos(salario))
}

func calcularImpuestos(salario float64) float64 {
	if salario > 150000.00 {
		return salario - salario*0.27
	} else if salario > 50000.00 {
		return salario - salario*0.17
	}
	return salario
}
