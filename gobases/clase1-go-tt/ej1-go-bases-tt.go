/*Ejercicio 1 - Letras de una palabra*/

package main

import "fmt"

func main() {
	text := "Academia"
	fmt.Printf("\n|| Letras de una palabra || \n\nPalabra designada: %s\n", text)
	fmt.Printf("> Cantidad de letras: %d \n", len(text))
	for _, index := range text {
		fmt.Print(" ", string(index))
	}
}
