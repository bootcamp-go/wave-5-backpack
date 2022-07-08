package main

import "fmt"

func main() {
	var palabra string = "Resumo"

	//fmt.Println(string(palabra[0]))

	fmt.Printf("Cantidad de letras: %d\n", len(palabra))
	for i := 1; i <= len(palabra); i++ {
		fmt.Printf("Letra %d: %c\n", i, palabra[i-1])
	}
}