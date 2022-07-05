package main

import (
	"fmt"
	"os"
	"strings"
)

func formatearArchivo(dir string) (string, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := os.ReadFile(dir)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	datos := strings.Split(string(data), ";")

	txt := fmt.Sprintf("ID \t\t\t Precio \t Cantidad\n")

	for i := 0; i < len(datos)-1; i++ {
		prod := strings.Split(datos[i], ",")
		txt += fmt.Sprintf("%s \t\t\t%s \t%s\n", prod[0], prod[1], prod[2])
	}

	return txt, err
}

func main() {
	formatearArchivo("dir")

	fmt.Println("Ejecucion completada!")
}
