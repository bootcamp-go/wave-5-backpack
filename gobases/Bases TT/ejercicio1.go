package main

import "fmt"
import "strings"

func main (){

	var palabra string = "holamundo"
	palabraDividida:= strings.Split(palabra,"")

	fmt.Printf("La cantidad de letras es: %v\n", len(palabraDividida))

	for  _, letter :=  range palabraDividida{
		fmt.Println(letter)
	}

	

}