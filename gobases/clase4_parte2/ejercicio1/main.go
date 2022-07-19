package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("./text.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Leyo el archivo de manera correcta %v", data)
}
