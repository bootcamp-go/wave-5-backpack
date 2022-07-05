package main

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return data
}

func main() {
	defer fmt.Println("ejecucion finalizada")
	data := ReadFile("customers.txt")
	fmt.Println(data)
}
