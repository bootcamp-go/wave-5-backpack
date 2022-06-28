package main

import (
	"fmt"
	"strconv"
)

func main() {

	palabra := "Hola"

	fmt.Printf("Numero de letras : %d\n", len(palabra))

	for _, letra := range palabra {
		fmt.Println(strconv.QuoteRune(letra))
	}
}
