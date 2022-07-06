// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

package main

import "fmt"

func impuestos(sueldo float32) float32 {
	if sueldo > 50000 {
		if sueldo > 150000 {
			return sueldo * 0.27
		} else {
			return sueldo * 0.17
		}
	} else {
		return 0
	}
}

func main() {
	var sueldo float32 = 300000
	fmt.Printf("El impuesto a descontar es: %.2f \n", impuestos(sueldo))
}
