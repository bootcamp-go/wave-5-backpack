package main

import (
	"fmt"
	"strings"
)

var palabra = "Palabra"

func main() {
	fmt.Println("La cantidad de letras que tiene la palabra es: ", len(palabra))
	palabraArreglo := strings.Split(palabra, "")
	for _, nombre := range palabraArreglo {
		fmt.Println(nombre)
	}
}
