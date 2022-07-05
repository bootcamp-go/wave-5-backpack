package main

import "fmt"

func calcularSalario(minutos int, categoria string) float64 {

	switch categoria {
	case "C":
		return 1000 * float64(minutos) / 60
	case "B":
		return 1.2 * 1500 * float64(minutos) / 60
	case "A":
		return 1.5 * 3000 * float64(minutos) / 60
	default:
		return 0
	}

}

func main() {

	var (
		salarioCategoriaA = calcularSalario(9600, "A")
		salarioCategoriaB = calcularSalario(9600, "B")
		salarioCategoriaC = calcularSalario(9600, "C")
	)

	fmt.Println("El salario base de la categoría A es:", salarioCategoriaA)
	fmt.Println("El salario base de la categoría B es:", salarioCategoriaB)
	fmt.Println("El salario base de la categoría C es:", salarioCategoriaC)

}
