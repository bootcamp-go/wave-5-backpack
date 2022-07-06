package main

import (
	"fmt"
	"os"
)

func leerTXT() {

	data, err := os.ReadFile("./custosmesrs.txt")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}

func main() {

	leerTXT()
	fmt.Println("Ejecución Finalizada")

}
