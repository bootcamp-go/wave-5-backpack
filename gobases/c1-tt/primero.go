package main

import "fmt"

func main() {
	var word string = "bootcamp"
	fmt.Printf("tiene %d palabras\n", len(word))
	for i := 0; i < len(word); i++ {
		fmt.Printf("letra: %q\n", word[i])
	}
}
