package main

import (
	"fmt"
	"log"
	"os"
)

type Producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func main() {
	productos := []Producto{
		{Id: 1, Precio: 2500, Cantidad: 4},
		{Id: 2, Precio: 5000, Cantidad: 6},
		{Id: 3, Precio: 100000, Cantidad: 10},
	}

	data := "Id;Precio;Cantidad\n"
	for _, value := range productos {
		line := fmt.Sprintf("%d;%.2f;%d;\n", value.Id, value.Precio, value.Cantidad)
		data += line
	}

	err := os.WriteFile("./productos.csv", []byte(data), 0644)
	if err != nil {
		log.Println(err)
	}
}
