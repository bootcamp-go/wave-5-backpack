package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./customers.txt"
	data := leerArchivo(fileName)
	fmt.Println(data)
}

func leerArchivo(fileName string) string {
	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	} else {
		return string(data)
	}
}
