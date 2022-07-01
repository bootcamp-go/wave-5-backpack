package main

import (
	"fmt"
)

const SALARIO_MINIMO = 150000

func main() {
	salary := 50000

	if salary < 150000 {
		fmt.Println(fmt.Errorf("error: el mínimo imponible es de %d y el salario ingresado es de %d", SALARIO_MINIMO, salary))
		return
	}

	fmt.Println("Debe pagar impuestos")
}
