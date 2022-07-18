package main

import (
	"fmt"
	"errors"
)

//Ejercicio 2 - Impuesto de salario #2

func main()  {
	
	salary := 160000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	} else {
		fmt.Println(errors.New("Debe pagar impuesto"))
		return
	}
}