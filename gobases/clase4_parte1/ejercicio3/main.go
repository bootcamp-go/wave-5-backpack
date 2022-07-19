package main

import (
	"fmt"
)

func main () {
	var salary int
	salary = 190000
	var err error
	if salary <= 0 {
		err = fmt.Errorf("el salario no puede ser menor o igual a 0")
		fmt.Println("ah ocurrido un error: ", err)
		return
	}
	if salary <= 150000 {
		err = fmt.Errorf("el salario ingresado no alcanza el mínimo imponible")
		fmt.Println("ah ocurrido un error: ", err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}