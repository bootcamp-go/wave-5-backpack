package main

import (
	"errors"
	"fmt"
	"os"
)

func myCustomError(value int) error {
	if value < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func main() {
	var salary int = 150000
	err := myCustomError(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
