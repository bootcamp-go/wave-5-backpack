package main

import (
	"fmt"
	"os"
)

func readCustomers() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	_, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado.")
	}
}

func main() {

	readCustomers()

	fmt.Println("ejecución finalizada")
}
