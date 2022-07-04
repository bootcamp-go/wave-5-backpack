package main

import "fmt"

const (
	A = "Categoria A"
	B = "Categoria B"
	c = "Categoria C"
)

func calcularSalario(categoria string, min float64) float64 {

	var horas float64
	var sueldoPorHora float64

	switch categoria {
	case "A":
		horas = min / 60
		sueldoPorHora = (horas * 3000) * 1.50
		return sueldoPorHora

	case "B":
		horas = min / 60
		sueldoPorHora = (horas * 1500) * 1.20
		return sueldoPorHora

	case "C":
		horas = min / 60
		sueldoPorHora = (horas * 1000) * 1.00
		return sueldoPorHora
	}
	return 0.0
}

func main() {
	fmt.Printf("El salario es: %v \n", calcularSalario("C", 60.0))

}
