package main

import (
	"fmt"
	"os"
)

func readFile(file string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := os.ReadFile(file)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func main() {
	readFile("customers.txt")
	fmt.Println("Ejecución Terminada")
}
