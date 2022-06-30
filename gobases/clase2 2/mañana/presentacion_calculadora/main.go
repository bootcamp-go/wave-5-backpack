package main

import (
	"fmt"
)

// Presentación -- Funciones en go
const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
)

// Suma
func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

// Resta
func opResta(valo1, valor2 float64) float64 {
	return valo1 - valor2
}

// Multiplicación
func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

// División
func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 { // validación división por cero
		return 0
	}

	return valor1 / valor2
}

func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorDeOperaciones(valores, opSuma)
	case Resta:
		return orquestadorDeOperaciones(valores, opResta)
	case Multip:
		return orquestadorDeOperaciones(valores, opMultip)
	case Divis:
		return orquestadorDeOperaciones(valores, opDivis)
	default:
		fmt.Printf("Operador %s no es válido", operador)
		return 0
	}
}

// Una función que recibe como argumento a otra función
func orquestadorDeOperaciones(valores []float64, operacion func(valor1, valor2 float64) float64) float64 {
	var resultado float64
	// iteramos los valores
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor) // Recursiva
		}
	}
	return resultado
}

func main() {
	fmt.Printf("Resultado %.0f\n", operacionAritmetica(Suma, 1, 2, 4, 5, 6)) // 18
	fmt.Printf("Resultado %.0f\n", operacionAritmetica(Resta, 1, 10))        // -9
	fmt.Printf("Resultado %.0f\n", operacionAritmetica(Multip, 1, 2))        // 2
	fmt.Printf("Resultado %.1f\n", operacionAritmetica(Divis, 1, 2))         // 0.5
}
