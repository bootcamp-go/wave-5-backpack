package main

import (
	"errors"
	"fmt"
	"os"
)

type producto struct {
	id       int
	precio   int
	cantidad int
}

func getCSV(productos ...producto) error {
	csvFormat := "ID,Precio,Cantidad"
	for _, p := range productos {
		newline := fmt.Sprintf("%d,%d,%d", p.id, p.precio, p.cantidad)
		csvFormat = fmt.Sprintf("%s\n%s", csvFormat, newline)
	}
	err := os.WriteFile("./myFile.csv", []byte(csvFormat), 0644)
	if err != nil {
		return errors.New("El archivo no pudo crearse")
	}
	return nil
}

func main() {
	producto1 := producto{1, 1000, 11}
	producto2 := producto{2, 20000, 1}
	producto3 := producto{3, 3300, 6}
	producto4 := producto{4, 4000, 2}

	getCSV(producto1, producto2, producto3, producto4)
}
