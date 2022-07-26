package main

import "fmt"

const (
	SUMA     string = "+"
	RESTA    string = "-"
	MULTIP   string = "*"
	DIVISION string = "/"
)

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case SUMA:
		return valor1 + valor2
	case RESTA:
		return valor1 - valor2
	case MULTIP:
		return valor1 * valor2
	case DIVISION:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

func main() {

	a, b := 6.5, 0.0

	fmt.Printf("SUMA: %.2f\n", operacionAritmetica(a, b, SUMA))
	fmt.Printf("RESTA: %.2f\n", operacionAritmetica(a, b, RESTA))
	fmt.Printf("MULTIP: %.2f\n", operacionAritmetica(a, b, MULTIP))
	fmt.Printf("DIVISION: %.2f\n", operacionAritmetica(a, b, DIVISION))

}
