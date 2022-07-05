package main

import "fmt"

func main() {
	word := "hola"

	fmt.Printf("%v \n", word)
	for i := range word {
		fmt.Printf("%v ", string(word[i]))
	}
	fmt.Println()
}
