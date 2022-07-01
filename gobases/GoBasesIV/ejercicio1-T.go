package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Vamos a leer el archivo...")

	str := "customer.txt"
	read(str)

	fmt.Println("Ejecucion finalizada")
}

func read(str string) {

	_, err := os.Open(str)

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
}
