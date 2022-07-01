package main

import (
	"fmt"
	"os"
)

func readFile(name string) *os.File {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	file, err := os.OpenFile(name, os.O_RDONLY, 0600)

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	return file
}

func main() {
	readFile("./archivo.txt")
	fmt.Println("Ejecución finalizada")
}
