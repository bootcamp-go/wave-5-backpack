package main

import "fmt"

func calcularSalario(categoria string, minutos float64) float64 {

	switch categoria {
	case "C":
		return (minutos / 60.0) * 1000.0
	case "B":
		res := (minutos / 60.0) * 1500.0
		return res + (res * 0.20)
	case "A":
		res := (minutos / 60.0) * 3000.0
		return res + (res * 0.60)
	}
	return 0

}

func main() {
	fmt.Printf("Sueldo : %2.f \n", calcularSalario("B", 60.0))
}
