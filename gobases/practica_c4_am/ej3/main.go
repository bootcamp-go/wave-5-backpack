package main

import "fmt"

func verificaciones(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}

func main() {

	var salary int = 10000000
	err := verificaciones(salary)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary = 15000
	err = verificaciones(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
