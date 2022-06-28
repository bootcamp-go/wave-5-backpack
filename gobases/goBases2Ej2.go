package main

import "fmt"

func main() {

	edad := 23
	sueldo := 110000
	antiguedad := 2

	if edad > 22 {
		if antiguedad > 1 {
			if sueldo > 100000 {
				fmt.Println("Puedes acceder al prestamo y no se te va a cobrar intereses.")
			} else {
				fmt.Println("Podes acceder al prestamo pero se cobra interes")
			}
		} else {
			fmt.Println("No cumplis con la antiguedad solicitada")
		}
	} else {
		fmt.Println("No tenes la edad suficiente")
	}
}
