package main

import "fmt"

func main(){
	i,j := 10,20
	j+=i
	fmt.Printf("resultado = %v \n",!(i != j))
}