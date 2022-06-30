package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")

	if err != nil {
		fmt.Printf("Error lectura directorio: %v", err)
	}

	fmt.Printf("files: %v \n", files)

	data, err := os.ReadFile("./paquete_fmt_1.go")

	if err != nil {
		fmt.Printf("Error lectura: %v", err)
	}

	fmt.Printf("file: %v \n", string(data))

	text := []byte("hello, gophers!")

	err = os.WriteFile("./myFile.txt", text, 0644)

	if err != nil {
		fmt.Printf("Error escritura: %v", err)
	}

}
