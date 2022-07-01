package main

import (
	"fmt"
)

func main(){
	var salary int = 140000

	//Ahora implementando Errorf para el msg con salario
	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de $150.000 y el salario ingresado es de: $%v", salary)
		fmt.Println(err)
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}