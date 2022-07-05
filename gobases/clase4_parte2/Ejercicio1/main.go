package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando ...")
	_, err := os.Open("customers.txt")

	if err != nil {
		fmt.Println("El archivo indicado no fue encontrado o está dañado")
		panic(err)

	}
	if err == nil || err != nil {
		fmt.Println("Ejecucion Finalizada")
	}
	fmt.Println("Ejecucion Finalizada")
}
