package main

/*
Guardar un archivo (write)
Datos de productos ()
*/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
	total    float64
}

func main() {
	/*prod1 := Producto{1, 15.00, 2, 0}
	prod2 := Producto{1, 11.40, 10, 0}

	dataProductos := []byte(formatData(prod1, prod2))
	err := os.WriteFile("./productos.csv", dataProductos, 0644)

	fmt.Println(formatData(prod1, prod2))
	if err != nil {
		fmt.Println("Error", err)
	}*/

	leerData()
}

func formatData(data ...Producto) string {
	var dataString string = "ID , Precio , Cantidad \n"
	for _, element := range data {
		dataString += fmt.Sprintf("%v , %v , %v \n", element.id, element.precio, element.cantidad)
	}

	return dataString
}

func leerData() {
	/*read, _ := os.ReadFile("productos.csv")
	data := fmt.Sprintf("%s", read)
	//fmt.Println(len(data))

	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	y := func(c rune) bool {
		return unicode.IsLetter(c)
	}
	fmt.Println(strings.FieldsFunc(data, x))
	fmt.Println(strings.FieldsFunc(data, y))

	//fmt.Printf(" %s", read)*/

	file, err := os.Open("productos.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	data := ""

	for {
		record, e := reader.Read()
		for key, e := range record {
			fmt.Println(data)

			if key == 0 {
				data += fmt.Sprintf("%s\t", e)
			} else {
				if key == (len(record))-1 {
					sFormated := e + strings.Repeat(" ", 9)
					data += fmt.Sprintf("%s\n", sFormated)
				} else {
					data += fmt.Sprintf("%s", e)

				}
			}
		}

		if e != nil {
			fmt.Println(e)
			break
		}
		//fmt.Println(record)
	}
}
