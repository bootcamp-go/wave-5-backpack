package main

import (
	"fmt"
	"strings"
)

var palabra string
var letras []string

func main() {

	palabra = "esternocleidomastoideo"
	letras = strings.Split(palabra, "")

	fmt.Println("La palabra tiene: ", len(letras), "letras")

	for _, letra := range letras {
		fmt.Println(letra)
	}

}
