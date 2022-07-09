package main

import "fmt"

var (
	impuesto1 float64 = 0.17
	impuesto2 float64 = 0.1
	tope1     float64 = 50000
	tope2     float64 = 150000
)

func calcularImpuesto(sueldo float64) (impuesto float64) {
	if sueldo > tope1 {
		impuesto += impuesto1
	}
	if sueldo > tope2 {
		impuesto += impuesto2
	}

	impuesto = sueldo * impuesto
	return
}

func main() {
	var sueldo float64

	fmt.Println("C2 GO Bases II - TM ")
	fmt.Println("Ingrese el sueldo: ")
	fmt.Scanln(&sueldo)
	fmt.Printf("El impuesto: %.2f \n", calcularImpuesto(sueldo))

}
