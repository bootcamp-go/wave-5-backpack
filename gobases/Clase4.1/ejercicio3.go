package main

import (
	"fmt"
)

func main() {
	salary := 100
	if salary < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de %v", salary)
		fmt.Println(err)
	}
}
