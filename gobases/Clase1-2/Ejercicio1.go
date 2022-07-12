package main

import "fmt"

func main(){
	var word string = "palabra"
	fmt.Printf("La palabra tiene %d caracteres\n",len(word))
	
	for _, s := range word{
		fmt.Printf("%c\n",s)
	}
}