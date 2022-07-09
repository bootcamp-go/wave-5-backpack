package main

import "fmt"

func main() {
	edad := 25
	empleado := true
	antiguedad_laboral := 2
	sueldo := 150000

	if edad > 22 {
		if empleado == true {
			if antiguedad_laboral > 1 {
				if sueldo > 100000 {
					fmt.Printf("Tu credito ha sido aprobado felicidades\n")
				} else {
					fmt.Printf("No cumples los requisitos, debes tener ingresos mayores a $100000\n")
				}
			} else {
				fmt.Printf("No cumples los requisitos, debes tener mas de 1 ano de antiguedad laboral\n")
			}
		} else {
			fmt.Printf("No cumples los requisitos, debes ser empleado\n")
		}
	} else {
		fmt.Printf("No cumples los requisitos, debes ser mayor de 22\n")
	}
}
