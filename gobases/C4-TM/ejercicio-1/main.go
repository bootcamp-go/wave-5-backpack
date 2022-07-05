package main

import "fmt"

type error1 struct {
}

func (err error1) Error() {
	fmt.Println("error: el salario ingresado no alcanza el mínimo imponible")
}

func main() {
	salary := 140000
	var myError error1

	if salary < 150000 {
		myError.Error()
	} else {
		fmt.Println("Debe pagar impuestos.")
	}
}
