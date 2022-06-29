package main

import "fmt"

func main() {
	var (
		nombre     = "Gabriel"
		edad       = 21
		empleado   = true
		antiguedad = 2
		salario    = 500.000
	)

	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Printf("cliente %s, se le otorga un prestamo ", nombre)
		if salario > 100.000 {
			fmt.Printf("con intereses\n")
		} else {
			fmt.Printf("sin intereses\n")
		}
	} else {
		fmt.Printf("cliente %s, no se le otorga un prestamo\n", nombre)
	}
}
