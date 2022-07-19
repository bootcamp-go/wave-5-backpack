package main

import (
	"fmt"
	"strings"
)

func main() {

	palabra := "Hola Mundo!"
	fmt.Println(len(palabra))

	var palabraSlice []string = strings.Split(palabra, "")

	for letra := range palabraSlice {
		fmt.Println(palabraSlice[letra])
	}
}
