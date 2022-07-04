package main

import (
	"fmt"
	"strings"
)

func main() {

	palabra := "Hello"

	fmt.Println("La longitud de palabra es: ", len(palabra))

	var letras []string = strings.Split(palabra, "")

	for i, letra := range letras {
		fmt.Println("Posición ", i, ": Letra", letra)
	}

}
