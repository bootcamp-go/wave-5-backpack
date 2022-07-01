package main

import "fmt"

type error interface {
	Error() string
}

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func main() {
	e := MyError{msg: "error: el salario ingresado no alcanza el minimo imponible"}
	salary1 := 50000

	if salary1 < 150000 {
		fmt.Println(e)
		return
	}
	fmt.Println("Debe pagar impuesto")
}
