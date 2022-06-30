package main

import "fmt"

const (
	CategoriaA = "A"
	CategoriaB = "B"
	CategoriaC = "C"
)

func calcularPorcAdicional(salario, porcAdicional float64) float64 {
	return salario + (salario * (porcAdicional / 100))
}

func calcularSalario(minutos, valorPorHora, porcAdicional float64) float64 {
	horas := minutos / 60
	salarioBase := horas * valorPorHora
	if porcAdicional > 0 {
		salarioBase = calcularPorcAdicional(salarioBase, porcAdicional)
	}
	return salarioBase
}

func calcularSalarioPorCategoria(minutos int, categoria string) float64 {
	switch categoria {
	case CategoriaA:
		return calcularSalario(float64(minutos), 3000, 50)
	case CategoriaB:
		return calcularSalario(float64(minutos), 1500, 20)
	case CategoriaC:
		return calcularSalario(float64(minutos), 1000, 0)
	default:
		return 0
	}
}

func main() {
	fmt.Println("Categoria A :", calcularSalarioPorCategoria(60, CategoriaA))
	fmt.Println("Categoria B :", calcularSalarioPorCategoria(60, CategoriaB))
	fmt.Println("Categoria C :", calcularSalarioPorCategoria(60, CategoriaC))
}
