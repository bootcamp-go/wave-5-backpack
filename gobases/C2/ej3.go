package main

import "fmt"

const (
	A = "A"
	B = "B"
	C = "C"
)

func calcularSalarios(cat string, minutos int) float64 {
	var horas float64

	horas = float64(minutos) / 60
	switch cat {
	case A:
		return float64((3000 * horas) + ((3000 * horas) * 0.5))

	case B:
		return float64((1500 * horas) + ((1500 * horas) * 0.2))

	case C:
		return float64(1000 * horas)
	}
	return 0

}

func main() {
	r := calcularSalarios(C, 60)
	fmt.Printf("%.2f \n", r)
}
