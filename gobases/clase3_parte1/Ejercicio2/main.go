package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../Ejercicio1/Productos.csv")
	if err != nil {
		fmt.Printf("Error lectura: %v", err)
	}
	r := string(data)
	fmt.Println(strings.ReplaceAll(r, ";", "\t\t\t"))

}
