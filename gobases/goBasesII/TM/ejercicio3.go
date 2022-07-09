package main

import (
	"fmt"
)

type categoria struct {
	sueldo float64
	bono   float64
}

var (
	catA categoria
	catB categoria
	catC categoria
)

func calSalario(temM int, categoria string) float64 {
	var (
		salario     float64
		tiempoHoras float64
	)
	tiempoHoras = float64(temM) / float64(60)
	if categoria == "a" || categoria == "A" {
		salario = (tiempoHoras * catA.sueldo * catA.bono)
	} else if categoria == "b" || categoria == "B" {
		salario = (tiempoHoras * catB.sueldo * catB.bono)
	} else if categoria == "c" || categoria == "C" {
		salario = (tiempoHoras * catC.sueldo * catC.bono)
	}
	return salario
}

func main() {
	catA = categoria{3000.0, 1.5}
	catB = categoria{1500.0, 1.2}
	catC = categoria{1000.0, 1.0}
	var (
		minutos   int
		categoria string
	)

	fmt.Println("Ingrese la cantidad de minutos trabajados: ")
	fmt.Scanln(&minutos)
	fmt.Println("Ingrese la categoria del trabajo: ")
	fmt.Scanln(&categoria)

	fmt.Printf("resultado: %.2f\n", calSalario(minutos, categoria))
}
