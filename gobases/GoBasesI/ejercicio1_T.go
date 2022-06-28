package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string = "mercadolibre"

	long := len(palabra)
	fmt.Println(long)
	fmt.Println("-------------------")

	letras := strings.Split(palabra, "")

	for _, letra := range letras {
		fmt.Println(letra)
	}
}
