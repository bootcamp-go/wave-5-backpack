package main

import "fmt"

func main() {

	var (
		edad       = 22
		empleado   = true
		antiguedad = 2
		sueldo     = 100
	)

	switch {
	case edad >= 22 && empleado && antiguedad > 1:
		if sueldo > 100000 {
			fmt.Println("Se otorga prestamo sin interes")
		} else {
			fmt.Println("Se otorga prestamo con interes")
		}
	default:
		fmt.Println("No se otorga prestamo")
	}

}
