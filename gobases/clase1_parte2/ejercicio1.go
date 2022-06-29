package main

import (
	"fmt"
	"strings"
)

func main() {

	var palabra = "Prueba"
	letras := strings.Split(palabra, "")
	g := ""

	fmt.Println("tamaño:", len(palabra))

	for i, letra := range letras {

		g = string(letra)
		fmt.Println("Letra:", g)
		i++
	}

}
