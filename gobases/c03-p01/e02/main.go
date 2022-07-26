package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	read, err := os.ReadFile("./e01/productos.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(string(read)) // formato byte
	datos := string(read)
	fmt.Println(strings.ReplaceAll(datos, ";", "\t\t\t"))

}
