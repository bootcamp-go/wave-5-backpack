package main

import (
	"fmt"
	"os"
)

type producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func main() {
	p1 := producto{111223, 30012.00, 1}
	p2 := producto{444321, 1000000.00, 4}
	p3 := producto{434321, 50.50, 1}

	str := fmt.Sprintf("%s;%s;%s\n%d;%.2f;%d\n%d;%.2f;%d\n%d;%.2f;%d", "ID", "Precio", "Cantidad", p1.ID, p1.Precio, p1.Cantidad, p2.ID, p2.Precio, p2.Cantidad, p3.ID, p3.Precio, p3.Cantidad)

	file := []byte(str)
	err := os.WriteFile("./informacion_producto.csv", file, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
