package main

import (
	"fmt"
	"os"
	"strings"
)

func leerArchivo(archivo string) string {
	data, err := os.ReadFile(archivo)

	if err != nil {
		return "Archivo no encontrado"
	} else {
		return string(data)
	}

}
func main() {

	//fmt.Println(leerArchivo("./productos.txt"))
	delimitador := "\n"
	productos := strings.Split(leerArchivo("./productos.txt"), delimitador)
	fmt.Print("ID\tPrecio\t\tCantidad\n")
	for _, linea := range productos {
		items := strings.Split(linea, ";")
		for _, item := range items {
			fmt.Print(item, "\t")
		}
		fmt.Println("")
	}

}
