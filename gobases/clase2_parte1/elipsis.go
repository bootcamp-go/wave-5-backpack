package main

import "fmt"

const (
	SUMA     string = "+"
	RESTA    string = "-"
	MULTIP   string = "*"
	DIVISION string = "/"
)

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

func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor)
		}
	}

	return resultado
}

func operacionAritmetica(operador string, valores ...float64) float64 {
	length := len(valores)
	fmt.Printf("len: %d\n", length)

	switch operador {
	case SUMA:
		return orquestadorOperaciones(valores, opSuma)
	case RESTA:
		return orquestadorOperaciones(valores, opResta)
	case MULTIP:
		return orquestadorOperaciones(valores, opMultip)
	case DIVISION:
		return orquestadorOperaciones(valores, opDivis)
	}

	return 0
}

func main() {

	r := operacionAritmetica(SUMA, 2, 3, 4, 7, 2, 7, 7, 4, 7, 8.6)

	fmt.Printf("resultado: %.2f, %.3f\n", r, r)

}
