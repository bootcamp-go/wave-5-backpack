package main

import "fmt"

var word string = "Camioneta"

func main(){
	fmt.Println("su palabra tiene: ", len(word))
	for _, letter := range word{
		fmt.Println("Letra: ", string(letter))
	}
}