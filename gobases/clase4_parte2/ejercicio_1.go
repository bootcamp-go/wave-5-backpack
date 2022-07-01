package main

import (
	"fmt"
	"os"
)

func leerArchivo(rutaArchivo string) string {
	data, err := os.ReadFile(rutaArchivo)

	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	} else {
		return string(data)
	}

}

func main() {
	rutaArchivo := "./customers.txt"
	leerArchivo(rutaArchivo)
}
