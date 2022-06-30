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

func print(archivo string) {
	delimitador := "\n"
	productos := strings.Split(archivo, delimitador)
	fmt.Print("ID\tPrecio\t\tCantidad\n")
	for _, linea := range productos {
		items := strings.Split(linea, ";")
		for _, item := range items {
			fmt.Print(item, "\t")
		}
		fmt.Println("")
	}
}
func main() {

	//fmt.Println(leerArchivo("./productos.txt"))
	print(leerArchivo("./productos.txt"))

}
