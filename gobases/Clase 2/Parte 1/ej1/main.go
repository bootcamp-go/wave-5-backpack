package main

import "fmt"

func calcularPorcentaje(num float64, porcentaje int64) float64 {
	return num * (float64(porcentaje) / 100)
}

func calcularDescuento(num float64, porcentaje int64) float64 {
	return num - calcularPorcentaje(num, porcentaje)
}

func descuentoSalarios(salario float64) float64 {
	var descuento int64
	if salario < 50000 {
		descuento = 0
	} else if salario < 100000 {
		descuento = 17
	} else {
		descuento = 27
	}

	return calcularDescuento(salario, descuento)
}

func main() {
	descuento := descuentoSalarios(50000)
	fmt.Println(descuento)
}
