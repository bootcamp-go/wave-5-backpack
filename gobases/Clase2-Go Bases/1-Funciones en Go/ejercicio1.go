package main

import "fmt"

func main(){
	salario:=500000.0
	fmt.Printf("Se debe descontar $%v \n",calcularImpuestos(salario)) 
}

func calcularImpuestos(sueldo float64) float64{
	var descuento float64=0.0
	if (sueldo>=50000){
		descuento=0.17*sueldo
		if(sueldo>=150000){
			descuento=descuento+(0.10*sueldo)
		}
	}
	return descuento
} 

/* 
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

*/