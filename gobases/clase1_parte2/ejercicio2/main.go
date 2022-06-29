package main

import "fmt"

func main() {

	var (
		edad       = 23
		trabaja    = true
		antiguedad = 2
		sueldo     = 50000
	)

	if edad > 22 && trabaja == true && antiguedad > 1 {

		if sueldo > 100000 {

			fmt.Println("Puede solicitar un crédito taza 0")

		}

		fmt.Println("puede solicitar un crédito a taza preferencial")

	}
	fmt.Println("por los momentos no cumple con los requisitos")

}
