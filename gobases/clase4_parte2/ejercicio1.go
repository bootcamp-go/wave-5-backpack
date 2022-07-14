package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Ejecución finalizada.")
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	read, err := os.ReadFile("./customer.txt")
	if err != nil {
		fmt.Println("Llegando al panic . . . 👀")
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	file := string(read)
	fmt.Println(file)
}
