package main

import (
	"fmt"
	"os"
	"strings"
)

func leerArchivo(archivo string) string {
	res, err := os.ReadFile(archivo)
	if err != nil {
		return "Hubo un error al leer el archivo"
	}
	return string(res)
}

func imprimirArchivo(archivo string){
	fmt.Printf("ID\t\t     Precio\t  Cantidad \n")
	for _, v := range archivo {
		data := strings.Split(string(v), ",")

		for i, value := range data {
			if i == 0 {
				fmt.Printf(value)
			} else if i == len(value)-1 {
				fmt.Printf("%v \n", value)
			} else {
				fmt.Printf("\t%v", value)
			}
		}
	}

	fmt.Println("\t  TOTAL: " )
}

func main() {
	imprimirArchivo(leerArchivo("../Ejercicio1/archivo.txt"))
}
