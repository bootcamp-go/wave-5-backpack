package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	productos := []producto{
		{1, 10.45, 3},
		{2, 15.80, 5},
		{3, 20, 10},
	}

	escribirCSV(productos)
}

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func escribirCSV(productos []producto) {
	data := fmt.Sprintln("ID,Precio,Cantidad")

	for i, p := range productos {
		if i == len(productos)-1 {
			data += fmt.Sprintf("%v,%v,%v", p.id, p.precio, p.cantidad)

			err := os.WriteFile("productos.csv", []byte(data), 0644)
			if err != nil {
				log.Fatal("No se puedo escribir el archivo")
			}
		}

		data += fmt.Sprintf("%v,%v,%v\n", p.id, p.precio, p.cantidad)
	}
}
