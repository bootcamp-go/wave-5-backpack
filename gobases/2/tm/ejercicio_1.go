package main

import (
	"errors"
	"fmt"
)

var s float64

func main() {
	s = 50001
	i, err := retencion(s)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Printf("El impuesto aplicado es de %.2f\n", i)
}

func retencion(salario float64) (impuesto float64, err error) {
	if salario < 0 {
		return 0, errors.New("No se admiten valores negativos")
	}

	impuesto = 0
	if salario > 50000 {
		impuesto = 0.17
	}
	if salario > 150000 {
		impuesto += 0.1
	}

	return salario * impuesto, nil
}
