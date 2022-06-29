package main

import "fmt"

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func main() {
	fmt.Println("Salario de la categoria A: ",calcSalario(600, categoriaA))
	fmt.Println("Salario de la categoria B: ",calcSalario(600, categoriaB))
	fmt.Println("Salario de la categoria C: ",calcSalario(600, categoriaC))
}

func calHoras(minutosTrabajados int) float64 {
	return float64(minutosTrabajados / 60)
}

func calcularAumento(horasTrabajadas, porcAdicional float64) float64 {
	return horasTrabajadas + (horasTrabajadas * (porcAdicional))
}

func calcSalario(minutosTrabajados int, categoria string) float64 {
	switch categoria {
	case categoriaA:
		return calcularAumento((calHoras(minutosTrabajados) * 3000), 0.5)
	case categoriaB:
		return calcularAumento((calHoras(minutosTrabajados) * 1500),0.2)
	case categoriaC:
		return float64(calHoras(minutosTrabajados)) * 1000
	default:
		return 0
	}
}
 
