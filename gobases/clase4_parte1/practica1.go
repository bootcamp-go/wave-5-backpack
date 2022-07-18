package main

import (
    "fmt"
    "os"
)

//Ejercicio 1 - Impuesto de salario #1

type IsError struct {
	message string
}

func (i *IsError) Error() string {
	return fmt.Sprintf("%v", i.message)
}

func IsErrorTest(salary int) error {
	if salary < 150000 {
		return &IsError {
			message : "error: el salario ingresado no alcanza el mínimo imponible" }
	} else {
		return &IsError {
			message : "Debe pagar impuesto"	}
	}
	return nil
	
}

func main()  {

	salary := 160000
	
	err := IsErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Funciona")
	
}