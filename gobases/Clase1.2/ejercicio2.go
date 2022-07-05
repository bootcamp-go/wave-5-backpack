package main

import "fmt"

func main() {
	age := 20
	pay := 500000

	if age <= 22 {
		fmt.Println("No apto para prestamo")
	} else {
		if pay < 100000 {
			fmt.Println("No se cobraran intereses")
		} else {
			fmt.Println("Se le cobraran intereses")
		}
	}
}
