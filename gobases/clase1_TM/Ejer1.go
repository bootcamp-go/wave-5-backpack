package main

import "fmt"

func main(){
	var Palabra = "donde estamos todos hoy"

	fmt.Println(len(Palabra)) // print length of string variable call Palabra
	for i:=0;i<len(Palabra);i++{
		fmt.Println(string(Palabra[i]))
	}


}