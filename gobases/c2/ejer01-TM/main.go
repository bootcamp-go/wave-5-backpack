package main

import "fmt"

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
// para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo
// y si gana más de $150.000 se le descontará además un 10%.

func calculoImpuesto(sueldo float32) float32 {

	var impuesto float32

	if sueldo > 50000 && sueldo < 150000 {
		impuesto = sueldo * 0.17
	} else if sueldo > 150000 {
		impuesto = sueldo * 0.27
	} else {
		impuesto = 0
	}

	return impuesto

}

func main() {

	fmt.Println(calculoImpuesto(51000))
}
