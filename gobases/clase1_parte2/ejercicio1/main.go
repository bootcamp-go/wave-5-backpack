package main

import "fmt"

func main() {
	var palabra = []string{"h", "o", "l", "a"}
	fmt.Println(len(palabra))

	for _, valor := range palabra {
		fmt.Println(valor)
	}
}
