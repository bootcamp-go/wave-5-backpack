package main

import "fmt"

func main() {
	palabra := "computadora"

	fmt.Printf("La palabra tieene %v letras \n", len(palabra))

	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}
