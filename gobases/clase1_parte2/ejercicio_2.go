package main

import "fmt"

func main() {
	edad := 23
	estado := true
	antiguedad_meses := 14
	sueldo := 200000

	if (edad > 22) && (estado == true) && (antiguedad_meses > 12) {
		fmt.Println("Se otorga el préstamo")
		if sueldo > 100000 {
			fmt.Println("No se cobrarán interéses")
		}
	} else {
		fmt.Println("No se otorga el préstamo")
	}
}
