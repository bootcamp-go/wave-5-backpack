package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 300000
	if salary < 150000 {
		fmt.Println(errors.New("error: no cumple con el salario minimo"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
