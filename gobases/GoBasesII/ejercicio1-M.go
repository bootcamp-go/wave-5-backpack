package main

import "fmt"

func main() {

	fmt.Println("Su salario con descuento es :", impuesto(150000))
}

func impuesto(salario float64) float64 {
	if salario > 150000 {
		return salario * (1 - 0.10)
	} else if salario > 50000 {
		return salario * (1 - 0.15)
	} else {
		return salario
	}
}
