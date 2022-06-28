package main

import "fmt"

func main() {
	palabra := "palabra"
	fmt.Printf("Numero de letras: %v \n", len(palabra))
	for _, letra := range palabra {
		fmt.Printf("letra: %v \n", string(letra))
	}
}
