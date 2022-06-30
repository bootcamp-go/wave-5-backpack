package main

/*
Guardar un archivo (write)
Datos de productos ()
*/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
	total    float64
}

func main() {
	prod1 := Producto{1, 15.00, 2, 0}
	prod2 := Producto{2, 11.40, 10, 0}
	prod3 := Producto{3, 12.50, 7, 0}

	dataProductos := []byte(formatData(prod1, prod2, prod3))
	err := os.WriteFile("./productos.csv", dataProductos, 0644)

	//fmt.Println(formatData(prod1, prod2))
	if err != nil {
		panic(err)
	}

	leerData()
}

func formatData(data ...Producto) string {
	var dataString string = "ID;Precio;Cantidad\n"
	for _, element := range data {
		dataString += fmt.Sprintf("%v;%v;%v\n", element.id, element.precio, element.cantidad)
	}

	return dataString
}

func leerData() {
	file, err := os.Open("productos.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	data := ""

	for {
		record, e := reader.Read()

		for key, e := range record {

			if key == 0 {
				data += fmt.Sprintf("%s\t\t", e)

			} else {
				if key == 2 {
					number, _ := strconv.ParseFloat(e, 32)
					if number > 0 {
						data += fmt.Sprintf("%8.2f\n", number)

					} else {
						fmt.Println("e", e)
						data += fmt.Sprintf("%s\n", e)

					}
				} else {
					data += fmt.Sprintf("%s\t", e)
				}
			}
		}

		if e != nil {
			fmt.Println(e)
			break
		}
	}

	fmt.Println(data)

}
