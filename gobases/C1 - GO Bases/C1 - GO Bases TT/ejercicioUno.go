package main

import(
	"fmt"
	"strings"
)

func main(){
	var word string = "Luz"
	//respuesta 1
	fmt.Println(len(word))

	//respuesta2
	fmt.Println("\n", strings.Split(word,""))

	//Alternativa a la respuesta 2
	for i:=0; i<len(word); i++{
		fmt.Println(string(word[i]))
	}
}