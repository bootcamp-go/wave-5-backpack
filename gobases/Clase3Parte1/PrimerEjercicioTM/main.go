package main

import (
	"fmt"
	"os"
)

func main() {
	prod := []byte("111111;5000;50\n111111;5000;50\n")
	err := os.WriteFile("./archivo.csv", prod, 0644)

	if err != nil {
		fmt.Printf("Hubo un error guardando el archivo")
	}
}
