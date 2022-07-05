package main

import (
	"fmt"
	"os"
)

type Productos struct {
	id       int
	precio   float64
	cantidad int
}

func pasarString(p []Productos) string {
	f := fmt.Sprintf("ID;\tPRECIO;\tCANTIDAD\n")
	total := 0.0
	for _, v := range p {
		total += (v.precio * float64(v.cantidad))
		f += fmt.Sprintf("%d;\t%.2f;\t%d\n", v.id, v.precio, v.cantidad)
	}

	f += fmt.Sprintf("Total:\t%.2f\n", total)
	return f

}
func main() {

	p := Productos{1, 10500, 4}
	z := Productos{2, 11500, 3}

	sl := []Productos{p, z}

	byt := []byte(pasarString(sl))

	err := os.WriteFile("./productos.csv", byt, 0644)

	if err != nil {
		fmt.Println("Error:", err) 

	}

	files, err2 := os.ReadFile("./productos.csv")

	if err2 != nil {
		fmt.Println("Error:", err2)

	} else {
		fmt.Printf(string(files))
	}

}
