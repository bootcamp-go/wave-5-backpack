package main

import (
	"errors"
	"fmt"
)

// Estructura y funcion que implementan el control de error
type salaryCheck struct {
	status int
	msg    string
}

func (e *salaryCheck) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

//------------------------------------------

func salaryTest(salary int) (int, error) {
	if salary < 150000 {
		return 500, &salaryCheck{
			status: salary,
			msg:    "El salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

// Control de error con errors.New()...super simple
func main() {
	var salary = 140000
	if salary < 150000 {
		fmt.Println(errors.New("El salario ingresado no alcanza el mínimo imponible"))
		return
	}
	//_, err := salaryTest(salary)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	fmt.Printf("%d - Debe pagar impuesto\n", salary)
}
