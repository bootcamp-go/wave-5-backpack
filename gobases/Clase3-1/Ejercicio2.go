package main

import (
	"fmt"
	"os"
)

func main(){
	data, _ := os.ReadFile("./productos.csv")
	txt:=string(data)
	fmt.Println(txt)
}