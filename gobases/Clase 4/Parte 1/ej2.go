package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 15000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}
