package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mostrarEnPantalla(archivo []byte) {
	inArchivo := string(archivo)

	inArchivo = strings.ReplaceAll(inArchivo, "ID: ", "")
	inArchivo = strings.ReplaceAll(inArchivo, "Precio: ", "")
	inArchivo = strings.ReplaceAll(inArchivo, "Cantidad: ", "")

	var datos []string = strings.Split(inArchivo, "\n")
	var productos []Producto
	var total float64

	for i := 0; i < len(datos)-1; i++ {
		var datosProducto []string = strings.Split(datos[i], "; ")
		datosProducto[i] = strings.ReplaceAll(datosProducto[i], " ", "")

		iD, err := strconv.ParseInt(datosProducto[0], 10, 64)
		precio, err := strconv.ParseFloat(datosProducto[1], 10)
		cantidad, err := strconv.ParseInt(datosProducto[2], 10, 64)

		if err != nil {
			fmt.Println(err)
		} else {
			productos = append(productos, newProducto(iD, precio, cantidad))
			//productos[i] = newProducto(iD, precio, cantidad)
		}
		fmt.Println(productos[i].imprimir(i))
	}

	for _, producto := range productos {
		total += producto.precio()
	}

	fmt.Printf("Total:%10.f", total)
}
