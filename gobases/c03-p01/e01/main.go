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

func total(productos []Producto) float64 {
	var precioTotal float64
	for _, p := range productos {
		precioTotal += float64(p.Precio) * float64(p.Cantidad)
	}
	return precioTotal
}

func main() {

	productos := []Producto{
		{Id: 1, Precio: 100, Cantidad: 10},
		{Id: 2, Precio: 200, Cantidad: 20},
		{Id: 3, Precio: 300, Cantidad: 30},
	}

	datos := "Id;Precio;Cantida\n"
	total := fmt.Sprintf(";%.2f;;", total(productos))

	for _, p := range productos {
		line := fmt.Sprintf("%d;%.2f;%d;\n", p.Id, p.Precio, p.Cantidad)
		datos += line
	}

	datos += total

	err := os.WriteFile("./productos.csv", []byte(datos), 0644)
	if err != nil {
		log.Println(err)
	}

}
