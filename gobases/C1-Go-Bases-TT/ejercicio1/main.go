package main

import "fmt"

// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
// 1. Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// 2. Luego imprimí cada una de las letras.

func main() {
	var palabra string = "Hola mundo mundial!"
	var totalLetras int = 0

	for _, v := range palabra {
		fmt.Printf("%c\n", v)
		totalLetras++
	}
	fmt.Printf("La palabra tiene %d letras", totalLetras)
}
