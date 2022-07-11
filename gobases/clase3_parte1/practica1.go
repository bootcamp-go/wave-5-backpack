package main

import (
	"fmt"
	"os"
	"log"
)

//Ejercicio 1 - Guardar archivo
type Producto struct {
	ID int
	Precio float64
	Cantidad int
}

func total(productos []Producto) float64 {
	var total float64
	for _, value := range productos {
		total += value.Precio * float64(value.Cantidad)
	}
	return total
}

func main()  {
	productos := []Producto {
		{ID: 1, Precio: 4500.00, Cantidad: 2},
		{ID: 2, Precio: 34200.40, Cantidad: 6},
		{ID: 3, Precio: 976400.00, Cantidad: 3},
	}

	data := "ID;Precio;Cantidad\n"
	total := fmt.Sprintf(";%.2f;;", total(productos))

	for _, value := range productos {
		row := fmt.Sprintf("%d;%.2f;%d;\n", value.ID, value.Precio, value.Cantidad)
		data += row
	}
	data += total 

	err := os.WriteFile("./productos.csv", []byte(data), 0644)
	if err != nil {
		log.Println(err)
	}
}