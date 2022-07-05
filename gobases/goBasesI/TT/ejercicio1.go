package main

import "fmt"

func main() {
	var palabra string
	fmt.Print("ingrese una palabra: ")
	fmt.Scanln(&palabra)
	fmt.Printf("Tamano: %d\n", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}
}
