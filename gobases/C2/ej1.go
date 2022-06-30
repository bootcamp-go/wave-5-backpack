package main

import "fmt"

func calcularImpuestos(sueldo float64) string {

	if sueldo < 50000 {
		return "Esta persona no paga impuestos"
	} else if sueldo > 50000 && sueldo < 150000 {
		var descuento float64
		descuento = 0.17

		impuesto := (sueldo * descuento)
		s := fmt.Sprintf("%.2f", impuesto)

		return "Esta persona debe pagar $ " + s + " de impuestos \n"
	} else {
		var descuento float64
		descuento = 0.27

		impuesto := (sueldo * descuento)
		s := fmt.Sprintf("%.2f", impuesto)

		return "Esta persona debe pagar $ " + s + " de impuestos \n"
	}
}

func main() {
	r := calcularImpuestos(151000)
	fmt.Printf(r)
}
