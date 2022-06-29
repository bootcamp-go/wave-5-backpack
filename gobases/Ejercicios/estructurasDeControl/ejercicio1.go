package main

import "fmt"

func letrasDeUnaPalabra() {
	//ejercicio 1
	palabra := "Hola"

	fmt.Printf("Cantidad de letras: %d \n", len(palabra))

	for i := range palabra {
		fmt.Printf("%c ", palabra[i])
	}

	fmt.Println()
}
