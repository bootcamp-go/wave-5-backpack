package main

import "fmt"

func calcularImpuesto(sueldo float64) float64{
	switch {
	case sueldo > 50000 && sueldo <= 150000:
		return sueldo*0.17
	case sueldo > 150000:
		return sueldo*0.10
	}
	return sueldo*0.2
}

var salario float64 = 130000

func main(){
	fmt.Println(calcularImpuesto(salario))
}