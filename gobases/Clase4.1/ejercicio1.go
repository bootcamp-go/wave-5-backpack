package main

import "fmt"

type myError struct {
}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el minimo imponible"
}

func main() {
	salary := 100
	if salary < 150000 {
		fmt.Println(&myError{})
	}
}
