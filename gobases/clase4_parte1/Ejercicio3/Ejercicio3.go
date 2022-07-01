package main

import (
	"fmt"
)

func main() {
	salary1 := 50000
	if salary1 < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de %d", salary1)
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}
