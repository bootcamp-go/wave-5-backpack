package main

import "fmt"

func impuestos(salario float64) float64 {
	var discount float64 = 0.00
	if salario > 50000.00 {
		discount = salario * 0.17
	}

	if salario > 150000.00 {
		discount = discount + (salario * 0.10)
	}

	return discount
}

func main() {
	var salario float64 = 100000.00
	fmt.Println(impuestos(salario))
}
