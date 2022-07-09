package main

import "fmt"

func main() {

    var palabra string

    palabra = "Hola Gente"
	fmt.Printf("ejercicio 1\n len %d\n", len(palabra))

    for i := 0; i < len(palabra); i++ {
        fmt.Printf("%c\n", palabra[i])
    }
}
