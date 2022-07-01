package main

import (
	"fmt"
)

//En mi estructura creo la variable para el mensaje
type MyError struct {
	msg    string
}

//Implementando Error()
func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

//Lanzando el mensaje de error
func messageError(salario int) (error) {
	return &MyError{
		msg:    fmt.Sprint("error: el salario ingresado: $",salario,", no alcanza el mínimo imponible"),
	}
}

func main(){
	var salary int = 140000
	err := messageError(salary)

	//Llamado al error de ser necesario
	if salary < 150000 {
		fmt.Println(err)
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}