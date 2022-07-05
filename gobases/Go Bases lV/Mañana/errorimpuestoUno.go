package main

import "fmt"

type ErrorSalario struct{}

func (e *ErrorSalario) Error() string {
	return "error: no cumple con el salario minimo"
}

func main() {
	salarioError := &ErrorSalario{}
	var salary int = 100000
	if salary < 150000 {
		fmt.Println(salarioError.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
