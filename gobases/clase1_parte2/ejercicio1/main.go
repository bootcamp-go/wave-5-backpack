package main

import (
	"fmt"
	"strings"
)

func main() {
	palabra := "Esto es una palabra"

	fmt.Println(len(palabra))

	separacion := strings.Fields(palabra)

	for _, separacion := range separacion {
		fmt.Printf("%s\n", separacion)
	}
}
