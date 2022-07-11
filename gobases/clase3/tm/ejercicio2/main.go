package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	byte, err := os.ReadFile("./../ejercicio1/productos.csv")
	if err != nil {
		log.Fatal(err)
	}

	data, err := leerArchivo(&byte)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func leerArchivo(data *[]byte) (string, error) { //Punteros pora practicar lo del turno tarde
	var s string

	var total float64

	lineas := strings.Split(string(*data), "\n")
	for i, linea := range lineas {
		fields := strings.Split(linea, ",")

		id := fields[0] // Campo ID en string

		// Set Cabecera
		if i == 0 {
			precio, cantidad := fields[1], fields[2] // Campo Precio y Cantidad en String

			s += fmt.Sprintf("%s\t\t%s\t%s\n", id, precio, cantidad)
			continue
		}

		// Agrega total en la última línea
		if i == len(lineas)-1 {
			s += fmt.Sprintf("\t\t%v", total) // Imprime total

			return s, nil
		}

		// Precio parseado a float64
		precio, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return "", err
		}

		// Cantidad parseado a float64
		cantidad, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return "", err
		}

		total += precio * cantidad

		// Agrega la fila con sus campos
		s += fmt.Sprintf("%v\t\t%v\t%v\n", id, precio, cantidad)
	}

	return s, nil
}
