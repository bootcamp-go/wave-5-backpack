package main

import "fmt"

func main() {
	var (
		edad uint = 99
		empleado bool = true
		antiguedad_en_anos float32 = 1.0
		salario float32 = 100000.1
	)

	//fmt.Println(!empleado && antiguedad_en_anos < 1)

	if edad <= 22 {
		fmt.Printf("El cliente debe ser mayor a 22 años\n")
	} else if !empleado || antiguedad_en_anos < 1 {
		fmt.Printf("El cliente debe estar empleado y tener más de 1 año de antigüedad\n")
	} else {
		if salario <= 100000 {
			fmt.Printf("El cliente puede optar a un préstamo y se le cobrará interés\n")
		} else {
			fmt.Printf("El cliente puede optar a un préstamo y no se le cobrará interés\n")
		}
	}
}