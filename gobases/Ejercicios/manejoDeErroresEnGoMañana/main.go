package main

import (
	"fmt"
	"os"
)

func main() {
	salary, err := impuestoSalaryError(100000)

	controlError(salary, err)

	salary, err = impuestoSalaryErrorNew(100000)

	controlError(salary, err)

	salary = 10000

	if salary < 150000 {
		fmt.Println(fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary))
		os.Exit(1)
	}

	fmt.Printf("Debe pagar impuestos. Su salario ($%d) supera los $150000\n", salary)

}
