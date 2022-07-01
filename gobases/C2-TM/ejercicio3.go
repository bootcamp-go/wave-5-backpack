package main

import "fmt"

func salario(minutosTrabajados int, categoria string) (salario float64) {
	horasTrabajadas := float64(minutosTrabajados) / 60
	switch categoria {
	case "A":
		salario = 3000 * horasTrabajadas * 1.5
	case "B":
		salario = 1500 * horasTrabajadas * 1.2
	case "C":
		salario = 1000 * horasTrabajadas
	}
	return
}

func main() {
	salario := salario(121, "C")
	fmt.Printf("El salario es: $%f\n", salario)
}
