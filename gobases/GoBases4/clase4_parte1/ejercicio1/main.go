package main

import (
	"fmt"
	"os"
)

func main() {

	var salary int = 100000

	salario, err := impuesto(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto por que su salario es %d \n", salario)
}

type mensualidad struct {
	salario int
	msg     string
}

func (m *mensualidad) Error() string {
	return fmt.Sprintf("%s ", m.msg)
}

func impuesto(salario int) (int, error) {
	if salario < 150000 {
		return 150000, &mensualidad{
			salario: salario,
			msg:     "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return salario, nil
}
