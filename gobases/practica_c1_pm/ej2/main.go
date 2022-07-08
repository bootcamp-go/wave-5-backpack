package main

import "fmt"

func main() {

	edad := 23
	empleado := true
	permanencia := 24
	salario := 100000

	switch {
	case edad > 22 && empleado && permanencia > 12:
		fmt.Println("Puede optar a Crédito")
		if salario > 100000 {
			fmt.Println("No se cobrará interés")
		} else {
			fmt.Println("Se cobrará interés")
		}
	default:
		fmt.Println("No puede otorgar crédito")
	}
}
