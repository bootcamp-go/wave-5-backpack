package main

import "fmt"

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {

	return fmt.Sprintf("%v", e.msg)
}
func validarImpuesto(salario int) error {
	if salario < 150000 {
		return &myCustomError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		return nil
	}
}

func main() {
	salary := 50000
	err := validarImpuesto(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}
