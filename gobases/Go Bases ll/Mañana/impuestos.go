package main

import "fmt"

var descuento float64

func calcularImpuesto(salario float64) float64 {
	if salario > 50000 && salario <= 150000 {
		descuento = 0.17
	} else if salario > 150000 {
		descuento = 0.1
	} else if salario < 50000 {
		descuento = 0
	}
	return descuento
}

func main() {
	calcularImpuesto(50000)
	fmt.Println(descuento)
}
