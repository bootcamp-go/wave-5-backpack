package main

import (
	"fmt"
)

const MIN_SALARY = 150000

func validarSalario(salary int) error {
	if salary < MIN_SALARY {
		return fmt.Errorf("error : el minimo imponible es de %d y el salario ingresado es de %d\n", MIN_SALARY, salary)
	}
	return nil
}

func main() {
	salary := 149999

	err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
	}

}
