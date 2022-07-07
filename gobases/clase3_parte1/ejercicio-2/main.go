package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Producto struct {
	id       int     `csv:"id"`
	precio   float64 `csv:"precio"`
	cantidad int     `csv:"cantidad"`
}

func leerProductos(nombreArchivo string) ([]Producto, error) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		return nil, errors.New("Error al leer el archivo")
	} else {
		var productos []Producto = []Producto{}
		datos := string(data)
		filas := strings.Split(datos, "\n")
		for i, fila := range filas {
			if i == 0 {
				// Saltando la cabecera del csv
				continue
			} else if fila == "" {
				// Saltando las filas vacias
				continue
			}

			fields := strings.Split(fila, ",")

			id, _ := strconv.Atoi(fields[0])
			precio, _ := strconv.ParseFloat(fields[1], 64)
			cantidad, _ := strconv.Atoi(fields[2])

			productos = append(productos, Producto{
				id:       id,
				precio:   precio,
				cantidad: cantidad,
			})
		}
		return productos, nil
	}
}

func imprimirProductos(productos []Producto) {
	padding := 3
	suma := 0.0
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight) //|tabwriter.Debug)
	fmt.Fprintln(w, " ID\tPrecio\tCantidad\t")
	for _, producto := range productos {
		fmt.Fprintf(w, "%d\t%.2f\t%d\t\n", producto.id, producto.precio, producto.cantidad)
		suma += producto.precio * float64(producto.cantidad)
	}
	fmt.Fprintf(w, "\t%.2f\t\n", suma)
	w.Flush()
}

func main() {
	productos, err := leerProductos("productos.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		imprimirProductos(productos)
	}
}
