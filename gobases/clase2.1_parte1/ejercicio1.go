package main

import "fmt"

//Ejercicio 1 - Impuestos del Salario

func impuesto(sueldo int) float64 {
	if sueldo > 50000 {
		return float64(sueldo) * 0.17
	} else if sueldo > 150000 {
		return float64(sueldo) * 0.27
	} else {
		return 0
	}
}

func main() {
	valImp := impuesto(500000)
	fmt.Println("El valor de su impuesto es: $", valImp)
}
