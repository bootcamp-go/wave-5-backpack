package main

import "fmt"

func main() {
	palabra := "Patricio"
	var longitud int = len(palabra)
	fmt.Printf("Cantidad de letras: %d \n", longitud)
	for i := 0; i < longitud; i++ {
		fmt.Printf("%c \n", palabra[i])
	}
}
