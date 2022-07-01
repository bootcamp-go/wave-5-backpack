package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}
	
func (err *myError) Error() string{
	return fmt.Sprintf("error: %s",err.msg)
}

func calcSalario(salario int) {
	if salario <  150000{
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}

func main(){
	salary := 160000
	calcSalario(salary)
}