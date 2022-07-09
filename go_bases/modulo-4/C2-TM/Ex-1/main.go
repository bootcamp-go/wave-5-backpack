package main

import "fmt"

func main() {
	var salario float32 = 155000
	descuento := calcSalario(salario)
	fmt.Println(descuento)
}
func calcSalario(salario float32) float32 {
	var impuesto float32

	if salario > 50 {
		impuesto = salario * 0.17
		fmt.Println("Impuesto del 17%")
	}
	if salario > 150000 {
		fmt.Println(impuesto)
		impuesto += (salario * 0.10)
		fmt.Println("+ 10% adicional")

	}

	return impuesto
}
