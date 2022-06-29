package main

import (
	"fmt"
)

const (
	catA     string = "A"
	catB     string = "B"
	catC     string = "C"
	salarioA int    = 1000
	salarioB int    = 1500
	salarioC int    = 3000
)

func main() {
	fmt.Println(calcularSalario(60, catA))
	fmt.Println(calcularSalario(90, catA))
	fmt.Println(calcularSalario(600, catB))
	fmt.Println(calcularSalario(60, catC))
}

func calcularSalario(minutos int, categoria string) (resultado float32) {
	horas := minutos / 60
	switch categoria {
	case catA:
		return float32(horas * salarioA)
	case catB:
		return float32(horas*salarioB) * 1.2
	case catC:
		return float32(horas*salarioC) * 1.5
	default:
		return -1
	}
}
