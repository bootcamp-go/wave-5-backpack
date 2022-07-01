package main

import "fmt"

const (
	CATEGORIA_A = "A"
	CATEGORIA_B = "B"
	CATEGORIA_C = "C"
)

const (
	SUELDO_BASE_A = 3000
	SUELDO_BASE_B = 1500
	SUELDO_BASE_C = 1000
)

const (
	ADICIONAL_A = 0.5
	ADICIONAL_B = 0.2
	ADICIONAL_C = 0
)

func main() {

	var salario = calcularSalario(60, CATEGORIA_C)

	fmt.Println(salario)
}

func calcularSalario(minutos int, categoria string) float32 {
	var cantHoras = float32(minutos) / 60

	switch categoria {
	case CATEGORIA_A:
		return cantHoras*SUELDO_BASE_A + (cantHoras*SUELDO_BASE_A)*ADICIONAL_A
	case CATEGORIA_B:
		return cantHoras*SUELDO_BASE_B + (cantHoras*SUELDO_BASE_B)*ADICIONAL_B
	case CATEGORIA_C:
		return cantHoras*SUELDO_BASE_C + (cantHoras*SUELDO_BASE_C)*ADICIONAL_C
	}

	return 0.0
}
