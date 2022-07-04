package main

import "fmt"

func main() {
	salary, err := impuestoSalary(1000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Debe pagar impuestos. Su salario ($%d) supera los $150000", salary)
	}

}
