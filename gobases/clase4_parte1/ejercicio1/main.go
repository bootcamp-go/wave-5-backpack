package main

import "fmt"

type customSalaryError struct {
	message string
}

func (se *customSalaryError) Error() string{
	return fmt.Sprintf("%s: ", se.message)
}

func validateSalary(salary int)(string, error){
	if salary <= 0 {
		return "", &customSalaryError{
			message: "el salario no puede ser menor o igual a 0",
		}
	}
	if salary < 150000 {
		return "", &customSalaryError{
			message: "el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return "Debe pagar impuesto", nil
}

func main () {
	var salary int
	salary = 90000
	message, err := validateSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

}