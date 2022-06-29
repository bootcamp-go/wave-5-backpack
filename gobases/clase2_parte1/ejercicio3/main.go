package main

import "fmt"

func main() {
	var cantMinutos float32 = 2000
	var categoria string = "A"
	sueldoFinal := salarioEmpleados(cantMinutos, categoria)
	fmt.Println("El sueldo del empleado final es: ", sueldoFinal)
}

func salarioEmpleados(cantMinutos float32, categoria string) float32 {
	var salario float32
	var horas float32 = cantMinutos / 60
	if categoria == "C" {
		return 1000.0 * horas
	}
	if categoria == "B" {
		salario = 1500.00 * horas
		return salario * 1.20
	}
	if categoria == "A" {
		salario = 3000.0 * horas
		return salario * 1.50
	} else {
		return salario
	}

}
