package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 160000

	status, err := salaryCheck(salary)
	if err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(status, "-Debe pagar impuesto 🔪")
}

func salaryCheck(s int) (int, error) {
	if s < 150000 {
		return 500, errors.New("❗️error: el salario ingresado no alcanza el mínimo imponible")
	}
	return 200, nil
}
