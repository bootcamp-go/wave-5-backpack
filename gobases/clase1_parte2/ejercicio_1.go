package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string
	palabra = "Hola"
	letras := strings.Split(palabra, "")
	fmt.Println("La palabra es", palabra, " y su longitud es de ", len(palabra))
	fmt.Println("Las letras por separado son: ", letras)
}
