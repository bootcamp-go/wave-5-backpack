package main

import "fmt"

const (
	Mult              = "*"
	categoriaA string = "A"
	categoriaB string = "B"
	categoriaC string = "C"
)

func cantidadSalario(minTrabajadosPorMes float64, categoria string) float64 {
	horas := minTrabajadosPorMes / 60
	switch categoria {
	case categoriaA:
		horasA := horas * 3000
		porcentajeAdicionalA := horasA * 0.50
		return horasA + porcentajeAdicionalA
	case categoriaB:
		horasB := horas * 1500
		porcentajeAdicionalB := horasB * 0.20
		return horasB + porcentajeAdicionalB
	case categoriaC:
		horasC := horas * 3000
		porcentajeAdicionalC := horasC * 0.50
		return horasC + porcentajeAdicionalC

	}
	return 0
}

func main() {
	fmt.Println("Salario A: ", (cantidadSalario(3000, categoriaA)))
	fmt.Println("Salario B: ", (cantidadSalario(3000, categoriaB)))
	fmt.Println("Salario C: ", (cantidadSalario(3000, categoriaC)))
}
