package main

import "fmt"

var word string = "kaleidoscope"

func main() {
	fmt.Printf("La cantidad de letras de la palabra %s es: %v \n", word, len(word))
	for i := 0; i < len(word); i++ {
		fmt.Printf("%q\n", word[i])
	}
}
