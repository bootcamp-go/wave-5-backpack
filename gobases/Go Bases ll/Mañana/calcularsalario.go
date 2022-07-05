package main

import "fmt"

func calcularSalario(categoria string, minutos float64) float64 {

	horasTrabajo := minutos / 60

	switch categoria {
	case "A":
		return (horasTrabajo * 3000) + (horasTrabajo*3000)*0.5

	case "B":
		return (horasTrabajo * 1500) + (horasTrabajo*3000)*0.2

	case "C":
		return (horasTrabajo * 1000)
	}
	return 0
}

func main() {
	fmt.Println("Salario mensual:", calcularSalario("D", 120))
}
