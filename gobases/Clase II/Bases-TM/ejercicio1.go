package main

import "fmt"

func impuestoSalario(sueldo float64) float64 {
	var impuesto float64
	if sueldo > 150000 {

		impuesto = (sueldo * 0.17) * 0.1

	} else if sueldo <= 150000 && sueldo > 50000 {
		impuesto = sueldo * 0.17

	} else {
		impuesto = 0.0
	}
	return impuesto
}

func main() {
	sueldo := 20000.0
	impuesto := impuestoSalario(sueldo)
	fmt.Printf("El impuesto es de: %v  \n", impuesto)

}
