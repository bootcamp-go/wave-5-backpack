package main

import "fmt"

func main() {
	var palabra string = "hola"

	fmt.Printf("tamaño: %d \n", len(palabra))

	for _, letra := range palabra {
		fmt.Printf("%q \n", string(letra))
	}

}
