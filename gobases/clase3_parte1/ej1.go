package main

import "fmt"

func main() {
	var sueldo = calcularImpuesto(100000)

	fmt.Println("El sueldo es: ", sueldo)
}

func calcularImpuesto(sueldo float32) float32 {
	if sueldo > 50000 && sueldo < 150000 {
		return sueldo * 0.83
	}

	if sueldo > 150000 {
		return sueldo * 0.73
	}

	return sueldo
}
