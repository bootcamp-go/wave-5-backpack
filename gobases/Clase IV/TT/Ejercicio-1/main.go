package main

import (
	"fmt"
	"os"
)

func leerTXT() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}

	}()

	data, err := os.ReadFile("./customers2.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(string(data))

}

func main() {

	leerTXT()
	fmt.Println("ejecución finalizada")

}
