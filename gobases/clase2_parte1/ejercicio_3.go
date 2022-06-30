package main

import "fmt"

func calcularSalario(minutos_trabajados float32, categoria string) float32 {
	horas_trabajadas := minutos_trabajados / 60

	switch categoria {
	case "C":
		salario := 1000 * horas_trabajadas
		return salario
	case "B":
		salario := 1500 * horas_trabajadas
		salario += salario * 0.2
		return salario
	case "A":
		salario := 3000 * horas_trabajadas
		salario += salario * 0.5
		return salario
	}
	return 0
}

func main() {
	fmt.Printf("El salario del empleado es %f\n", calcularSalario(14400.0, "A"))
}
