package main

import (
	"fmt"
	"os"
)

type Productos struct {
	id       string
	precio   float64
	cantidad int
}

func toString(productos []Productos) string {
	var count float64
	p := fmt.Sprintf("ID\tPRECIO\tCANTIDAD\n")

	for _, producto := range productos {
		count += producto.precio * float64(producto.cantidad)
		p += fmt.Sprintf("%s\t%.2f\t%d\n", producto.id, producto.precio, producto.cantidad)
	}
	p += fmt.Sprintf("Total:\t%.2f", count)
	return p
}

func main() {
	// Creando archivo

	var p1 Productos = Productos{"1", 3.5, 4}
	var p2 Productos = Productos{"2", 3.5, 4}
	var p3 Productos = Productos{"3", 3.5, 4}

	productos := []Productos{p1, p2, p3}

	createCSV := []byte(toString(productos))

	err := os.WriteFile("./productos.csv", createCSV, 0644)

	if err != nil {
		fmt.Println("Hubo un error:", err)
	}

	// leyendo archivo
	files, err := os.ReadFile("./productos.csv")

	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
	fmt.Println(string(files))
}
