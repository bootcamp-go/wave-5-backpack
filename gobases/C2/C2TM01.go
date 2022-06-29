package main

import (
	"fmt"
)

func calcularImpuesto(sueldo float64)  float64{
	impuesto := 0.0

	if sueldo > 50000{
		if sueldo > 150000{
			impuesto= sueldo * 0.17
		} else {
			impuesto= sueldo * 0.27
		}
	} 

	return impuesto
}
func main() {

	sueldo := 600000.0
	fmt.Println("El impuesto es: ", calcularImpuesto(sueldo))

}