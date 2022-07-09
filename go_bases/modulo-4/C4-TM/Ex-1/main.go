package main

import (
	"fmt"
)

type salaryError struct {
	status int
	msg    string
}

func (e *salaryError) Error() string {
	return fmt.Sprintf("%d-%v", e.status, e.msg)
}

func main() {

	/*
		En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
		Crea un error personalizado con un struct que implemente “Error()”
		con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
		Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
	*/

	var salary int = 160000

	status, err := salaryCheck(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status, "-Debe pagar impuesto 🔪")

}

func salaryCheck(s int) (int, error) {
	if s < 150000 {
		return salaryStatus(500)
	}
	return salaryStatus(200)
}
func salaryStatus(s int) (int, error) {
	if s == 500 {
		return 500, &salaryError{
			status: s,
			msg:    "❗️error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}
