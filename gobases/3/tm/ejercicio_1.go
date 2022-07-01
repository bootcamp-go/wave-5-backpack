package main

import (
	"fmt"
	"os"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	productos := []producto{
		{1, 12.0, 5},
		{2, 3.50, 10},
		{3, 2.25, 7},
		{4, 1.75, 22},
	}

	var rows string = "id,precio,cantidad\n"
	for _, v := range productos {
		rows += fmt.Sprintf("%d,%.2f,%d\n", v.id, v.precio, v.cantidad)
	}
	os.WriteFile("./gobases/3/tm/ejercicio_1.csv", []byte(rows), 0644)
}
