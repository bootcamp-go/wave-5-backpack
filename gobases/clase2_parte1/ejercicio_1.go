package main

import (
	"fmt"
)

func calcularImpuesto(salario float64) float64 {
	if salario > 50000 && salario <= 150000 {
		salario -= salario * 0.17
	} else if salario > 150000 {
		salario -= salario * 0.27
	} else {
		salario = 0
	}
	return salario
}

func main() {
	salario := 160000.0
	fmt.Println("El salario del empleado es ", salario, " y el impuesto de su salario es ", calcularImpuesto(salario))
}
