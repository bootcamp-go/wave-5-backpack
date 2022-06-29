package main

import "fmt"

const (
	CATEGORIA_C = "categoria_c"
	CATEGORIA_B = "categoria_b"
	CATEGORIA_A = "categoria_a"
)

func calcularPorcentage(numero int, porcentage int) int {
	return int(float32(numero) * (float32(porcentage) / float32(100)))
}

func transformarMinutosAHora(minutos int) float32 {
	return float32(minutos) / float32(60)
}

func calcularSalario(minutos int, categoria string) int {
	switch categoria {
	case CATEGORIA_C:
		sueldoMes := transformarMinutosAHora(minutos) * 1000
		return int(sueldoMes)
	case CATEGORIA_B:
		sueldoMes := int(transformarMinutosAHora(minutos) * 1500)
		return sueldoMes + calcularPorcentage(sueldoMes, 20)
	case CATEGORIA_A:
		sueldoMes := int(transformarMinutosAHora(minutos) * 1500)
		return sueldoMes + calcularPorcentage(sueldoMes, 50)
	}
	return 0
}

func printSalario(minutos int, tipo string) {
	fmt.Println("Salario Trabajado", minutos, "min para empleado", tipo, ":", calcularSalario(minutos, tipo))
}

func main() {
	minutosTrabajados := 120
	printSalario(minutosTrabajados, CATEGORIA_C)
	printSalario(minutosTrabajados, CATEGORIA_B)
	printSalario(minutosTrabajados, CATEGORIA_A)

}
