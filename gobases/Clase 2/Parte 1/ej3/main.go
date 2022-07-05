package main

import "fmt"

func calcularPorcentaje(num float64, porcentaje int64) float64 {
	return num * (float64(porcentaje) / 100)
}

func calcularAumento(num float64, porcentaje int64) float64 {
	return num + calcularPorcentaje(num, porcentaje)
}

func calcularSalario(categoria string, minTrabajados uint) float64 {
	horasTrabajadas := minTrabajados / 60
	var aumento int64
	var salarioHora float64 = 1000
	if categoria == "B" {
		aumento = 20
		salarioHora = 1500

	} else if categoria == "A" {
		aumento = 50
		salarioHora = 3000
	}
	return calcularAumento(salarioHora*float64(horasTrabajadas), aumento)
}

func main() {
	sueldoTotal := calcularSalario("B", 120)
	fmt.Println(sueldoTotal)
}
