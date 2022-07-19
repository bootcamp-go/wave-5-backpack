package main

import (
	"errors"
	"fmt"
)

func main () {
	var salary int
	salary = 90000
	if salary <= 0 {
		fmt.Println(errors.New("el salario no puede ser menor o igual a 0"))
		return
	}
	if salary <= 150000 {
		fmt.Println(errors.New("el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}