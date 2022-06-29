package main

import "fmt"

func main() {
	var palabra = "Computadora"
	fmt.Println(len(palabra))

	for _, valor := range palabra {
		fmt.Println(string(valor))
	}
}
