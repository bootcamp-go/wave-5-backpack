package main

import (
	"fmt"
)

func main() {
	var salary int = 90000
	if salary < 150000 {
		err := fmt.Errorf("error: el salario minimo es %d", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
