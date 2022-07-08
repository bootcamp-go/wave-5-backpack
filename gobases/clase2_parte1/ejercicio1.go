package main

import "fmt"

func imp_salario(salario float64) float64 {
	
	if salario > 50000.00 {
		return salario * 0.17
	} else if salario > 100000.00 {
		return salario * 0.27
	}

	return 0
}

func main() {
	
	impuesto := imp_salario(51000.00)

	if impuesto != 0 {
		fmt.Printf("El impuesto aplicable al empleado es de: %f\n", impuesto)
	} else {
		fmt.Println("El empleado no posee impuesto aplicable")
	}
}