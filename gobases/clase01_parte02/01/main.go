package main

import "fmt"

var palabra string = "hola"

func main(){
	fmt.Println("Cantidad de letras: ",len(palabra))
	for i, letra := range palabra {
		fmt.Println("letra", i,":", string(letra))
	}
}