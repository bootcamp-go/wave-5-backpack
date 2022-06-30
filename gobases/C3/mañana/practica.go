package main

import (
	"fmt"
	"os"
)

func main() {
	text := []byte("1123321;502.51;3\n1123333;230;2")
	err := os.WriteFile("./productos.csv", text, 0644)

	if err != nil {
		fmt.Printf("Error escritura: %v", err)
	}

}
