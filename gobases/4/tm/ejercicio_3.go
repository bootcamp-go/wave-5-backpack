package main

import (
	"fmt"
	"os"
)

func myCustomError(value int) error {
	if value < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", value)
	}
	return nil
}

func main() {
	var salary int = 15000
	err := myCustomError(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
