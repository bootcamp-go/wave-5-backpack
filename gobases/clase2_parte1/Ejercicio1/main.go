package main

import "fmt"

func main(){
	fmt.Println(impuestoSalario(50.000))
}

func impuestoSalario(salario float64) float64{
	if salario > 50000 && salario < 150000 {
		return salario - 0.17
	}else{
		return salario - 0.27
	}
}