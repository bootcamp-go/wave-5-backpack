package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	leerArchivo("./customers.txt")

	fmt.Println("ejecución finalizada")
}

func leerArchivo(nombre string) []byte {
	defer func() {
		err := recover()

		if err != nil {
			log.Printf("panic al leer el archivo: %v\n", err)
		}
	}()

	data, err := os.ReadFile(nombre)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	return data
}
