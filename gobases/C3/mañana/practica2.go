package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./productos.csv")

	if err != nil {
		fmt.Printf("Error lectura: %v", err)
	}
	r := string(data)

	var out string = "\n"
	var boolean bool
	for _, letra := range r {

		if string(letra) == ";" {
			if !boolean {
				boolean = true
				out = out + "\t \t"
			}
			out = out + " "
		} else if letra == 10 {
			boolean = false
			out = out + "\n"
		} else {
			out = out + string(letra)
		}

	}
	fmt.Println("ID  \t \t Precio Cantidad \n", out)

}
