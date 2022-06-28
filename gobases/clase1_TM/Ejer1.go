package main

import "fmt"

func main(){
	var Palabra = "donde estamos todos hoy"

	fmt.Println(len(Palabra)) // print length of string variable call Palabra
	for i:=0;i<len(Palabra);i++{
		fmt.Println(string(Palabra[i]))
	}

	for _, v :=range Palabra{ // using for range , where de index is not used change for character _ to omit 
		fmt.Printf(v)
	}

}