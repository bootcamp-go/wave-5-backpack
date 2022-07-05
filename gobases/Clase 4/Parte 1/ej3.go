package main

import "fmt"

func main() {
	var salary int = 15000

	if salary < 150000 {
		fmt.Println(fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary))
		return
	}
	fmt.Println("Debe pagar impuesto")
}
