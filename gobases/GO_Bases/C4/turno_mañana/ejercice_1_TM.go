package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	myError := &MyError{}
	salary := 200000
	if salary < 150000 {
		fmt.Println(myError.Error())
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
