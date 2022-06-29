package main

import "fmt"

func obtenerDescuento(salario, porcentaje float64) (float64, float64) {
	descuento := salario * (porcentaje / 100)
	return descuento, (salario - descuento)
}

func calcularImpuesto(salario float64) float64 {
	totalSalario := salario
	totalDescuentos := 0.0
	if salario > 50000 {
		totalDescuentos, totalSalario = obtenerDescuento(salario, 17)
	}
	if salario > 150000 {
		desc, _ := obtenerDescuento(totalSalario, 10)
		totalDescuentos += desc
	}
	return totalDescuentos
}

func main() {
	fmt.Println("Valor final :", calcularImpuesto(200000))
}
