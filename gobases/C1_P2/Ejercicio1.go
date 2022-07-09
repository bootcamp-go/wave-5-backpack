package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string = "World"
	w := strings.Split(word, "")
	fmt.Printf("%q", w)
	fmt.Println("La palabra tiene", len(word), "caracteres")
	for i := 0; i < len(w); i++ {
		fmt.Println(w[i])
	}
}
