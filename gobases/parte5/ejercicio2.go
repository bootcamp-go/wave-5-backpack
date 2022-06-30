package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("./productosComprados.csv")
	pos := 0
	total := .0
	values := []string{}
	fmt.Printf("ID\t\t    Precio   Cantidad\n")
	for i, b := range data {
		if b == ';' || b == '\n' {
			values = append(values, string(data[pos:i]))
			pos = i + 2
		}
		if b == '\n' {
			pos--
			id, _ := strconv.Atoi(values[0])
			precio, _ := strconv.ParseFloat(values[1], 64)
			cantidad, _ := strconv.Atoi(values[2])
			fmt.Printf("%d\t\t%10.2f   %8d\n", id, precio, cantidad)
			total += float64(cantidad) * precio
			values = []string{}
		}
	}
	fmt.Printf("\t\t%10.2f\n", total)
}
