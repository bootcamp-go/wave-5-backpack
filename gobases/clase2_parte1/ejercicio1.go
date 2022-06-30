/*
	Ejercicio 1 - Impuestos de salario
	
	Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de 
	depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva 
	el impuesto de un salario. 
	
	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo 
	y si gana más de $150.000 se le descontará además un 10%.
*/

package main

import "fmt"

func main() {
	sueldo := 160000.0
	fmt.Println("El impuesto es ", calculoImpuesto(sueldo))
}

func calculoImpuesto(sueldo float64) float64 {
	var impuesto float64 = 0
	if sueldo > 50000 && sueldo < 150000 {
		impuesto = sueldo * 0.17
	} else if sueldo > 150000 {
		impuesto = sueldo * 0.27
	}
	return impuesto
}

