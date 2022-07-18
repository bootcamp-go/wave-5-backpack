package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func main() {
	var productos []producto
	data, err := os.ReadFile("./myFile.csv")
	fmt.Println(err)
	dataString := string(data)
	dataList := strings.Split(dataString, "\n")
	for _, d := range dataList {
		if d != "" {
			dSplit := strings.Split(d, ",")
			id, _ := strconv.Atoi(dSplit[0])
			precio, _ := strconv.ParseFloat(dSplit[1], 64)
			cantidad, _ := strconv.Atoi(dSplit[2])
			productos = append(productos, producto{id, precio, cantidad})
		}
	}
	fmt.Printf("%s\t%20s%10s\n", "ID", "Precio", "Cantidad")
	var (
		total         float64
		cantidadTotal int
	)

	for _, p := range productos {
		fmt.Printf("%d\t%20.2f\t%6d\n", p.ID, p.Precio, p.Cantidad)
		total += (p.Precio * float64(p.Cantidad))
		cantidadTotal += p.Cantidad
	}
	fmt.Printf("\t%20.2f\t%6d\n", total, cantidadTotal)
}
