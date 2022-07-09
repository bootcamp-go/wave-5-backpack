package main

import (
	"fmt"
	"os"
)

type producto struct {
	ID       int
	Precio   float32
	Cantidad int
}

var productos []producto

func main() {
	p1 := producto{1000, 5000.0, 2}
	p2 := producto{1001, 5000.0, 3}
	p3 := producto{1002, 5000.0, 4}
	p4 := producto{1003, 5000.0, 5}
	p5 := producto{1004, 5000.0, 6}
	p6 := producto{1005, 5000.0, 7}
	p7 := producto{1006, 5000.0, 8}

	productos = append(productos, p1)
	productos = append(productos, p2)
	productos = append(productos, p3)
	productos = append(productos, p4)
	productos = append(productos, p5)
	productos = append(productos, p6)
	productos = append(productos, p7)

	var data string
	for _, p := range productos {
		data += fmt.Sprintf("%d,%.2f,%d\n", p.ID, p.Precio, p.Cantidad)
	}
	dataBit := []byte(data)
	err := os.WriteFile("./myFile.csv", dataBit, 0644)
	fmt.Print(err)
}
