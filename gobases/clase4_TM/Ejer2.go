package main

import (
	"errors"
	"fmt"
)

func testSalary(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: eres muy pobre")

		//		"error: el salario ingresado no alcanza el mínimo imponible"
	}
	return "debe pagar impuesto", nil
}

func main() {
	var salary int
	fmt.Println("ingrese el salario")
	fmt.Scanln(&salary)

	msg, err := testSalary(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", msg)

}
