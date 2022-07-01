package main

// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

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

// La misma empresa necesita leer el archivo almacenado, para ello requiere que:
// se imprima por pantalla mostrando los valores tabulados,
// con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
// el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

// Ejemplo:

// ID                            Precio  Cantidad
// 111223                      30012.00         1
// 444321                    1000000.00         4
// 434321                         50.50         1
//                           4030062.50

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
	fmt.Println("ID\t\t\t    Precio\t\t\tCantidad")
	for _, prod := range l.Productos {
		fmt.Printf("%s\t\t\t%10.2f\t\t\t%8d\n", prod.Id, prod.Precio, prod.Cantidad)
	}
	fmt.Printf("Total\t\t\t%10.2f\t\t\t\n", l.calcularTotal())
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
		"1", 100, 2,
	}

	prod2 := Producto{
		"2", 200, 10,
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
