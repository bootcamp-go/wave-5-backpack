package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListaProductos struct {
	Productos []Producto
}

type Producto struct {
	Id       string
	Precio   float64
	Cantidad int
}

func (l ListaProductos) toCSV() string {
	str := "Id;Precio;Cantidad\n"
	for _, prod := range l.Productos {
		str += fmt.Sprintf("%s;%.2f;%d\n", prod.Id, prod.Precio, prod.Cantidad)
	}
	return str
}

func (l ListaProductos) writeAsCSV() error {
	btw := []byte(l.toCSV())
	return os.WriteFile("./products.csv", btw, 0644)
}

func (l *ListaProductos) readFromCSV(data []byte) {
	fileStr := string(data)
	campos := strings.Split(fileStr, "\n")
	for i, linea := range campos {
		if i != 0 && linea != "" {
			variables := strings.Split(linea, ";")
			if len(variables) == 3 {
				id := variables[0]
				precio, err := strconv.ParseFloat(variables[1], 64)
				cantidad, err2 := strconv.Atoi(variables[2])
				if err != nil && err2 != nil {
					fmt.Println("Error parseando datos en linea", i)

				} else {
					p := Producto{
						Id:       id,
						Precio:   precio,
						Cantidad: cantidad,
					}
					l.Productos = append(l.Productos, p)
				}
			}
		}
	}
}

func (l ListaProductos) print() {
	fmt.Println("ID\t    Precio\tCantidad")
	for _, prod := range l.Productos {
		fmt.Printf("%s\t%10.2f\t%8d\n", prod.Id, prod.Precio, prod.Cantidad)
	}
	fmt.Printf("Total\t%10.2f\t\n", l.calcularTotal())
}

func (l ListaProductos) calcularTotal() float64 {
	total := 0.0
	for _, prod := range l.Productos {
		total += prod.Precio * float64(prod.Cantidad)
	}
	return total
}

func main() {
	prod1 := Producto{
		Id:       "1",
		Precio:   100,
		Cantidad: 2,
	}

	prod2 := Producto{
		Id:       "2",
		Precio:   200,
		Cantidad: 10,
	}

	prod3 := Producto{
		Id:       "3",
		Precio:   200,
		Cantidad: 3,
	}

	prodList := []Producto{prod1, prod2, prod3}
	lista := ListaProductos{prodList}

	err := lista.writeAsCSV()

	if err != nil {
		fmt.Println("Error escribiendo archivo")
	}

	file, err2 := os.ReadFile("./products.csv")
	if err2 != nil {
		fmt.Println("Error leyendo archivo")
	}

	listaLectura := ListaProductos{}
	listaLectura.readFromCSV(file)
	listaLectura.print()
}
