package main

import (
	"fmt"
	"os"
)

func leerTXT() {

	data, err := os.ReadFile("./customerss.txt")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
		fmt.Println("ejecución finalizada")

	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}

func main() {

	leerTXT()

}
