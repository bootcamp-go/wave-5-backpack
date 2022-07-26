package main

import "fmt"

func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func operacionAritmetica(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Suma":
		return opSuma
	case "Resta":
		return opResta
	case "Multip":
		return opMultip
	case "Divis":
		return opDivis
	}
	return nil
}

func main() {
	oper := operacionAritmetica("No suma")
	r := oper(2, 5)
	fmt.Printf("resultado = %.2f\n", r)
}
