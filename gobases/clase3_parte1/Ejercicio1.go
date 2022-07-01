package main

import (
	"fmt"
	"os"
)

func main() {

	texto := "ID,PRECIO,CANTIDAD\n1,100,20,\n2,250,25,\n3,500,40"
	txt := []byte(texto)
	err := os.WriteFile("./Archivo.txt", txt, 0644)
	if err != nil {
		fmt.Println("Ocurrio un error.")
	}
}
