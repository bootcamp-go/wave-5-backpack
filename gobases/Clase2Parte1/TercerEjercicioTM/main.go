package main

import "fmt"

func salarioHorasTrabajadas(minutos float64, porcentaje float64, sueldo float64) float64 {
	horas := minutos / 60
	sueldoHoras := sueldo * horas
	return sueldoHoras + sueldoHoras*porcentaje/100
}

func salarioTotal(categoria string, minutos float64) float64 {
	switch categoria {
	case "A":
		return salarioHorasTrabajadas(minutos, 50, 3000)
	case "B":
		return salarioHorasTrabajadas(minutos, 20, 1500)
	case "C":
		return salarioHorasTrabajadas(minutos, 0, 1000)
	default:
		return 0
	}
}
func main() {
	salarioEmpleado := salarioTotal("C", 25000)
	fmt.Printf("El salario del empleado es: %v \n", salarioEmpleado)
}
