package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg string
}
	
func (err *myError) Error() string{
	return fmt.Sprintf("error: %s",err.msg)
}

func calcSalario(salario int) (string,error)  {
	if salario < 150000 {
		return "", &myError { msg: "el salario ingresado no alcanza el mínimo imponible"}
	}
	return "Debe pagar impuesto", nil
}

func main(){
	salary := 150500
	response, err := calcSalario(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response)
}