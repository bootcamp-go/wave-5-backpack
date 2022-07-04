package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string = "pikachu"
	fmt.Println(len(palabra))

	letras := strings.Split(palabra, "")

	for _, letra := range letras {
		fmt.Println(letra)
	}

}
