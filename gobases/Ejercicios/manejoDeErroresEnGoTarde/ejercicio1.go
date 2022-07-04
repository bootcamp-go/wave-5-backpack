package main

import (
	"fmt"
	"os"
)

func leerArchivo(name string) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.ReadFile(name)

	if err != nil {
		panic("El archivo no existe")
	} else {
		fmt.Println(file)
	}

}
