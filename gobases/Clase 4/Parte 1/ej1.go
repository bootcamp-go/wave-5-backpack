package main

import "fmt"

type CustomErr struct{}

func (c CustomErr) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	var salary int = 1500000

	if salary < 150000 {
		err := CustomErr{}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Debe pagar impuesto")
}
