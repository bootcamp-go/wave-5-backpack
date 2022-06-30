package main

import (
	"fmt"
)

func main(){

}

func calcularImpuesto(salario int)float64{
	if salario > 50000 {
		result := (salario * 100) / 17
		fmt.Println("Tu salario es de", result)
	}else if salario > 150000 {
		result := (salario * 100) / 27
		fmt.Println("Tu salario es de", result)
	}
	return 0
}