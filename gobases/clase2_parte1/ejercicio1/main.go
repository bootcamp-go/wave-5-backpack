package main

import (
	"errors"
	"fmt"
)

func main() {
	sueldo := 180000
	impuesto, err := calcularImpuestoxSalario(float32(sueldo))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El impuesto para el sueldo: ", sueldo, " es: ", impuesto)
	}
}
func calcularImpuestoxSalario(sueldo float32) (float32, error) {
	if sueldo < 0 {
		return 0, errors.New("El salario debe ser un numero positivo")
	}

	impuesto := 0.0
	if sueldo > 50000 {
		impuesto += .17
	}
	if sueldo > 150000 {
		impuesto += .10
	}
	return (sueldo * float32(impuesto)), nil
}
