package main

import (
	"fmt"
	"os"
	"strings"
)

type Productos struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	prod1 := Productos{111223, 30012.00, 1}
	prod2 := Productos{444321, 1000000.00, 4}

	str := fmt.Sprintf("%d, \t %.2f, \t %d \n%d, \t %.2f, \t %d \n", prod1.id, prod1.precio, prod1.cantidad, prod2.id, prod2.precio, prod2.cantidad)
	err := os.WriteFile("./archivo.txt", []byte(str), 0644)

	if err != nil {
		fmt.Println(err)
	}

	/* PARTE DOS */

	read, err := os.ReadFile("./archivo.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("ID\t\tPrecio\t\tCantidad \n")
	for _, v := range read {
		data := strings.Split(string(v), ",")

		for i, value := range data {
			if i == 0 {
				fmt.Printf(value)
			} else if i == len(value)-1 {
				fmt.Printf("%v \n", value)

			} else {
				fmt.Printf("\t %v", value)
			}
		}
	}
	fmt.Println("total: " )

	//fmt.Printf("ID \t Precio \t Cantidad \n%v", string(read))
}
