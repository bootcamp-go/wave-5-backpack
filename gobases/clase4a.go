package main

import (
	"errors"
	"fmt"
	"os"
)

type myerror struct {
	status int
	msg    string
}

func (e *myerror) Error() string {
	return fmt.Sprintf("%d status, mensaje: %s\n", e.status, e.msg)
}
func custum_salary_error(salary int) (string, error) {
	if salary < 150000 {
		return "error:el salario no alcanza el minimo perdedor\n", &myerror{
			status: 404,
			msg:    "buena la cagaste pobre ql\n",
		}
	}
	return "cuico ql paga la wea deve impuesto\n", nil
}
func custum_salary_error2(salary int) (string, error) {
	if salary < 150000 {
		return "error:el salario no alcanza el minimo perdedor\n",
			errors.New("404 status, mensaje: buena la cagaste pobre ql\n")
	}
	return "cuico ql paga la wea deve impuesto\n", nil
}
func custum_salary_error3(salary int) (string, error) {
	if salary < 150000 {
		return "error:el salario no alcanza el minimo perdedor\n",
			fmt.Errorf("error %d: el mínimo imponible es de 150.000 y el salario ingresado es de: %d\n", 404, salary)
	}
	return "cuico ql paga la wea deve impuesto\n", nil
}
func main() {
	fmt.Println("polo")

	var salary int
	salary = 27400
	mesage, err := custum_salary_error3(salary)
	if err != nil {
		fmt.Println(err, mesage)
		os.Exit(1)
	}
	fmt.Println(mesage)

}
