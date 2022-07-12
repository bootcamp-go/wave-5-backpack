package main

import "fmt"

func main(){
	//var salario float64 = 160000
	var salario float64 = 50000
	var impuesto float64 = CalculoImpuesto(salario)
	fmt.Printf("el impuesto es %.2f \n",impuesto)
}

func CalculoImpuesto(salario float64) float64{
	var impuesto float64
	if salario > 150000 {
		impuesto = (27 * salario)/100
	} else if salario > 50000{
		impuesto = (17 * salario)/100
	} else {
		impuesto = 0
	}
	return impuesto
}