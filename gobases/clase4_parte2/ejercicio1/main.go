package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ejecucion finalizada")
	}()
	_, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo no fue encontrado o esta dañado")
	}
}
