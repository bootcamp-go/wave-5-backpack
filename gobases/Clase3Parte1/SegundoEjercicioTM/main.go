package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./archivo.csv")
	fmt.Printf("Id\tPrecio\tCantidad\n")
	if err != nil {
		fmt.Printf("Hubo un error en la lectura")
	} else {
		for _, value := range data {
			val := strings.Split(string(value), ";")
			for i, v := range val {
				if i == 0 {
					fmt.Printf(v)
				} else if i == len(v)-1 {
					fmt.Printf("\t %v \n", v)
				} else {
					fmt.Printf("\t %v", v)
				}
			}
		}
	}
}
