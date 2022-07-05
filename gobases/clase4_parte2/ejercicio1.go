package main

import (
	"fmt"
	"os"
)

//Ejercicio 1 - Datos de Clientes

func main() {
	var pathFile string
	pathFile = "/a.txt"
	//readFile(pathFile)

	defer func(path string) ([]byte, error) {
		data, err := os.ReadFile(pathFile)
		if err != nil {
			panic("el archivo indicado no fue encontrado o está dañado")
		}
		return data, err
	}(pathFile)
	fmt.Println("Ejecucion Finalizada")
}
