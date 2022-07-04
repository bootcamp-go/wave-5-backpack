package main

import "fmt"

func main() {

	var (
		edad       int     = 28
		empleado   bool    = true
		antiguedad int     = 3
		salario    float32 = 180000
	)

	if edad > 22 && empleado && antiguedad > 1 {
		if salario > 100000 {
			fmt.Println("Puede acceder al crédito sin interés")
		} else {
			fmt.Println("Puede acceder al crédito con interés")
		}
	} else {
		fmt.Println("No puede acceder al crédito")
	}

}
