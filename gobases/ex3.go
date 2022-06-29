package main

import (
	"fmt"
	"strings"
)

func ex3() {
	words := "Hola como va?"
	word := strings.Split(words, "")
	fmt.Println("cantidad de caracteres:", len(word))
	for i := 0; i < len(word); i++ {
		fmt.Print(word[i], " \n")
	}
}
