package main

import (
	"fmt"
	"os"
	"strings"
)

//Ejercicio 2 - Leer archivo
func main()  {
	readFile, err := os.ReadFile("./productos.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(readFile)
	fmt.Println(strings.ReplaceAll(data, ";", "\t\t\t"))
}