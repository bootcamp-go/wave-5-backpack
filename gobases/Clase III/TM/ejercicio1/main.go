package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	Precio   float64
	Cantidad int
}

func exportarCSV(values []Producto) string {

	var text string
	text = "id,precio,cantidad\n"
	for _, value := range values {
		text += fmt.Sprintf("%d,%0.2f,%d\n", value.id, value.Precio, value.Cantidad)
	}
	data := []byte(text)
	err := os.WriteFile("./archivo.csv", data, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return text
}

func main() {

	ProductosAll := []Producto{{1, 30012.00, 1}, {2, 1000000.00, 4}, {3, 50.50, 1}}
	productos := exportarCSV(ProductosAll)

	fmt.Println(productos)

}
