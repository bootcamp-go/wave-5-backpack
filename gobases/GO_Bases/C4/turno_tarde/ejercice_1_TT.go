package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Ejecución finalizada.")
		err := recover() // Recupera el panic - para evitar una ejecución no deseada

		if err != nil {
			fmt.Println(err)
		}
	}()

	read, err := os.ReadFile("./customer.txt")
	if err != nil {
		fmt.Println("PANIC!")
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	file := string(read)
	fmt.Println(file)
}
