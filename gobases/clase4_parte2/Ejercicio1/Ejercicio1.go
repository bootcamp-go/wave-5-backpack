package main

import (
	"fmt"
	"os"
)

func leerArchivo(path string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := os.ReadFile(path)
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	fmt.Println("Archivo leido con exito")
}

func main() {
	leerArchivo("hola.txt")
	leerArchivo("hola2.txt")
	fmt.Println("ejecucion finalizada")
}
