package main

import (
	"errors"
	"fmt"
)

func main() {
	salary1 := 50000
	if salary1 < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}
