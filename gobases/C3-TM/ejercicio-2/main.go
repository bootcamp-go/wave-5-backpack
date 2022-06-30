package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prt, err := formatearArchivo("./prod_comprados.txt")

	if err != nil {
		fmt.Println("Hubo un problema al leer el archivo.")
	} else {
		fmt.Print(prt)
	}
}

func formatearArchivo(dir string) (string, error) {
	data, err := os.ReadFile(dir)
	datos := strings.Split(string(data), ";")

	txt := fmt.Sprintf("ID \t\t\t Precio \t Cantidad\n")

	for i := 0; i < len(datos)-1; i++ {
		prod := strings.Split(datos[i], ",")
		txt += fmt.Sprintf("%s \t\t\t%s \t%s\n", prod[0], prod[1], prod[2])
	}

	return txt, err
}
