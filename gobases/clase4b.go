package main

import (
	//"context"
	"fmt"
	"os"
)

var hola int = 4

func readFile(path string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo estaba corrupto")
			return
		}
	}()
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func main() { //"customer.txt"
	// fmt.Println("jpa")
	// fmt.Println(hola)
	// hola = 5
	// fmt.Println(hola)
	// hola = 6
	texto := readFile("nada.txt")
	println(texto)
	// defer func() {
	// 	err := recover()
	// 	fmt.Println("mira tu")
	// 	if err != nil {
	// 		fmt.Println("pinche ", err)
	// 	}
	// 	hola = 7
	// 	fmt.Println(hola)
	// }()
	// defer func() {
	// 	hola = 7
	// 	fmt.Println(hola)
	// 	fmt.Println("ayayaya")
	// }()
	// if hola == 6 {
	// 	fmt.Println("hihi")
	// 	panic("carajo")
	// }
	fmt.Println("todo finish")
}
