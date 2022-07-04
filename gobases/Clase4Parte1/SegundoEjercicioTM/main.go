package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 1500000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
