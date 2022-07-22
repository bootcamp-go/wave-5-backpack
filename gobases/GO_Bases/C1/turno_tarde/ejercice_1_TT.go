package main

import (
	"fmt"
)

func main() {
	palabra := "hola"
	fmt.Println("el largo de la palabra es: %v", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
