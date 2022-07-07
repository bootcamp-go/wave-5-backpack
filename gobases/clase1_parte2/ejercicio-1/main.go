package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		La Real Academia Española quiere saber cuántas letras tiene una palabra
		y luego tener cada una de las letras por separado para deletrearla.
			1. Crear una aplicación que tenga una variable con la palabra e
			imprimir la cantidad de letras que tiene la misma.
			2. Luego imprimí cada una de las letras.
	*/

	// 1.
	var palabra string = "Hola"
	fmt.Println(len(palabra))

	// 2.
	var palabraSlice []string = strings.Split(palabra, "")

	for letra := range palabraSlice {
		fmt.Println(palabraSlice[letra])
	}
}
