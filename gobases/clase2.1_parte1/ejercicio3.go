package main

import (
	"fmt"
)

//Ejercicio 3 - Calcular Salario

func calSalario(categoria string, minutos int) float64 {
	switch categoria {
	case "C":
		return float64((minutos / 60) * 1000)
	case "B":
		sal := float64((minutos / 60) * 1500)
		sal += sal * 0.20
		return sal
	case "A":
		sal := float64((minutos / 60) * 3000)
		sal += sal * 0.50
		return sal
	}
	return 0
}

func main() {
	category := "A"
	minutes := 60
	fmt.Println("Su salario para la Categoria", category, "es $",
		calSalario(category, minutes), "Habiendo trabajado ", minutes, "Minutos")
}
