package main

import (
	"fmt"
	"strings"
)

func main() {
	palabra := "edificio"
	delimitador := ""
	letras := strings.Split(palabra, delimitador)
	fmt.Println("Palabra: ", palabra)
	fmt.Println("Numero de letras: ", len(letras))
	for _, letra := range letras {
		fmt.Println(letra)
	}
}