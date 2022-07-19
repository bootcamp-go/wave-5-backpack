package main

import (
	"fmt"
	"os"
)

func leerArchivo(archivo string) {
	file, err := os.ReadFile(archivo)
	if err != nil {
		panic("No existe archivo")
	}

	fmt.Println(string(file))
}

func main() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Proceso finalizado")
	}()

	leerArchivo("customers.txt")
}
