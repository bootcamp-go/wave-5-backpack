package main

import "fmt"

func calcularImpuesto(salario float64) float64 {
	if salario > 50000 {
		if salario > 150000 {
			return 0.27 * salario
		}
		return 0.17 * salario
	}
	return 0
}

func main() {
	impuesto := calcularImpuesto(200000)
	fmt.Println("El impuesto del salario es:", impuesto)
}
