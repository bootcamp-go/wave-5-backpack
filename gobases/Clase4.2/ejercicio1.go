package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Ejecucion finalizada.")
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile("./customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}
