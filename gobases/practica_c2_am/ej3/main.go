package main

import "fmt"

func main() {
	minutosTrabajados := 200
	categoria := "A"
	salario := calculoSalarioEmpleado(minutosTrabajados, categoria)
	fmt.Println("El salario es: ", salario)

	minutosTrabajados = 200
	categoria = "B"
	salario = calculoSalarioEmpleado(minutosTrabajados, categoria)
	fmt.Println("El salario es: ", salario)

	minutosTrabajados = 200
	categoria = "C"
	salario = calculoSalarioEmpleado(minutosTrabajados, categoria)
	fmt.Println("El salario es: ", salario)
}

func calculoSalarioEmpleado(cantiMinutosTrabajados int, categoria string) float64 {
	switch categoria {
	case "A":
		return (float64(cantiMinutosTrabajados) * 3000) * 1.5
	case "B":
		return (float64(cantiMinutosTrabajados) * 1500) * 1.2
	case "C":
		return float64(cantiMinutosTrabajados) * 1000
	default:
		return 0
	}
}
