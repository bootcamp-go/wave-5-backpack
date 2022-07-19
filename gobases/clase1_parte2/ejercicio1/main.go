package main

import "fmt"

func main() {
	var word = "Inyeccion"
	fmt.Println("Tu palabra tiene: ", len(word))
	for _, letter := range word {
		fmt.Println("Tu apalabra elegida es: ", string(letter))
	}
}
