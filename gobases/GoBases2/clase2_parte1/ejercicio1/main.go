package main

import "fmt"

func calcularImpuesto(salario float64) float64 {
	switch {
	case salario > 150000:
		return ((salario - (salario * 0.15)) * 0.1) + (salario * 0.15)
	case salario > 50000 || salario < 150000:
		return salario * 0.15
	default:
		return 0
	}

}

func main() {

	fmt.Println(calcularImpuesto(72000))
	fmt.Println(calcularImpuesto(40000))
	fmt.Println(calcularImpuesto(160000))

}
