package main

import "fmt"

// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
// Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// Luego imprimí cada una de las letras.

func main() {
	var palabra = "Stefano"
	var letras []string

	for i := 0; i < len(palabra); i++ {
		letras = append(letras, string(palabra[i]))
	}

	fmt.Println("Cantidad de letras de la palabara:", len(palabra))
	for i := 0; i < len(palabra); i++ {
		println(letras[i])
	}
}
