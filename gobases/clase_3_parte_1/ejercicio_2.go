package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Producto struct {
	id       int
	precio   int
	cantidad int
}

func (p *Producto) setPropiedad(propiedad string, valor string) {
	switch propiedad {
	case "id":
		dato, err := strconv.ParseInt(valor, 10, 32)
		if err == nil {
			p.id = int(dato)
		}
	case "precio":
		dato, err := strconv.ParseInt(valor, 10, 32)
		if err == nil {
			p.precio = int(dato)
		}
	case "cantidad":
		dato, err := strconv.ParseInt(valor, 10, 32)
		if err == nil {
			p.cantidad = int(dato)
		}
	}
}

func leerProductoCSV(filePath string, separador string) ([]Producto, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	csvProduct := strings.Split(string(data), "\n")
	propiedades := []string{}

	productos := []Producto{}
	for i, line := range csvProduct {
		if i == 0 {
			propiedades = strings.Split(line, separador)
		} else {
			producto := Producto{}
			valores := strings.Split(line, separador)
			if len(valores[0]) > 0 {
				for i, valor := range valores {
					producto.setPropiedad(strings.Trim(propiedades[i], " "), strings.Trim(valor, " "))
				}
				productos = append(productos, producto)
			}

		}

	}
	return productos, nil
}

func printProducts(products []Producto) {
	fmt.Printf("%.5s \t %10s \t %10s \n", "ID", "Precio", "Cantidad")
	for _, product := range products {
		fmt.Printf("%.5d \t %10d \t %10d \n", product.id, product.precio, product.cantidad)
	}
}

func main() {
	products, err := leerProductoCSV("./productos.csv", ";")
	if err != nil {
		fmt.Println(err)
	} else {
		printProducts(products)
	}
}
