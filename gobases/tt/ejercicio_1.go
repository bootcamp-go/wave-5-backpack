package main

import "fmt"

func main() {
	word := "gobases"
	fmt.Printf("Chars counted: %d\n", len(word))

	for i, _ := range word {
		fmt.Printf("%c\n", word[i])
	}
}