package main

import "fmt"

func calcImpuesto(salario float64) float64 {
	if salario > 150000 {
		return salario * 0.27
	} else if salario > 50000 {
		return salario * 0.17
	} else {
		return 0
	}
}

func main() {
	sueldo := 1000000
	fmt.Printf("Se le descuenta: %v\n", calcImpuesto(float64(sueldo)))
}
