package main

import (
	"fmt"
)

func calcSalario(salario int) {
	if salario <  150000{
		fmt.Println(fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salario))
		return
	}
	fmt.Println("Debe pagar impuesto")
}

func main(){
	salary := 140000
	calcSalario(salary)
}