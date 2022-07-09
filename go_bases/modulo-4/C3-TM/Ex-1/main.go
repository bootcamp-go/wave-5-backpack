package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var myShop [3][]string
	var ID, Precio, Cantidad []string
	ID = append(ID, "ID")
	Precio = append(Precio, "Precio")
	Cantidad = append(Cantidad, "Cantidad")

	myShop[0] = ID
	myShop[1] = Precio
	myShop[2] = Cantidad

	fmt.Println(myShop)

	err2 := os.WriteFile("Gopher.csv", []byte("Hello Go %t Precio%t Producto"), 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	data, err := os.ReadFile("Gopher.csv")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
