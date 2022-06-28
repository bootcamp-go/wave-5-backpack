package main

import "fmt"

func main() {
	var palabra = "Camilo"
	var lPalabra = len(palabra)

	fmt.Printf("La longitud de %s es de %v caracteres \n", palabra, lPalabra)

	for idx, letra := range palabra {
		fmt.Printf("Letra %d: %c \n", idx+1, letra)
	}

}
