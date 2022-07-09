package main

import "fmt"

//error: el salario ingresado no alcanza el mínimo imponible
//Debe pagar impuesto

const topeSalario float64 = 150000

type MyError struct {
	arg  float64
	prob string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func impuestoSalario(arg float64) (float64, error) {
	if arg < topeSalario {
		return -1, &MyError{arg, "el salario ingresado no alcanza el mínimo imponible"}
	}
	return arg, nil
}

func main() {
	var salary float64

	fmt.Scanln(&salary)

	if _, e := impuestoSalario(salary); e != nil {
		fmt.Println("error: ", e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	_, e := impuestoSalario(salary)
	if ae, ok := e.(*MyError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
