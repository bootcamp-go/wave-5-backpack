package main

import "fmt"

type salaryError struct {
	salary int
}

func (se *salaryError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return &salaryError{salary}
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main() {
	salary := 146000

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
	}
}
