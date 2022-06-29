package main

import "fmt"

func descuentoSalario(salario float64) float64 {
	if salario > 49999 && salario < 150000 {
		return salario - ((salario * 17) / 100)
	} else if salario > 149999 {
		return salario - ((salario * 27) / 100)
	} else {
		return salario
	}
}
func main() {
	salarioTotal := descuentoSalario(75000)
	fmt.Printf("El salario total es: %v \n", salarioTotal)
}
