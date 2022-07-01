package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	productosLeidos := leerArchivo()
	imprimirArchivo(productosLeidos)
}

func leerArchivo() []string {
	bytesLeidos, _ := os.ReadFile("./myFile.txt")
	return strings.Split(string(bytesLeidos), "\n")
}

func imprimirArchivo(productosLeidos []string) {
	cabecera := "ID \t\t Precio \t Cantidad"
	fmt.Println(cabecera)

	var total float64 = 0
	for _, producto := range productosLeidos {
		productoSeparado := strings.Split(producto, ",")
		if len(productoSeparado) < 3 {
			continue
		}
		registroAImprimir := fmt.Sprint(productoSeparado[0], "\t\t", productoSeparado[1], "\t\t", productoSeparado[2], "\t")
		fmt.Println(registroAImprimir)
		precio, _ := strconv.ParseFloat(strings.TrimSpace(productoSeparado[1]), 64)
		cantidad, _ := strconv.ParseFloat(strings.TrimSpace(productoSeparado[2]), 64)
		total += precio * cantidad
	}
	fmt.Println("\t\t", total)
}
