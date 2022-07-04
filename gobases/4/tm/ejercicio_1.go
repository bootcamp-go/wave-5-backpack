package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomError(value int) error {
	if value < 150000 {
		return &myError{msg: "error: el salario ingresado no alcanza el mínimo imponible"}
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
