package main

import (
	"fmt"
)

func main() {
	edad := 25
	tiempo_empleado:=1
	sueldo:= 20000

	if edad > 22{

		if tiempo_empleado > 1{

			if sueldo > 100000{

				fmt.Println("Credito aprobado y se cobra interes")

			}else {
				fmt.Println("Credito parobado y no cobra interes")
			}

		} else {
			fmt.Println("No cumple la edad minima")
		}

	} else {
		fmt.Println("No cumple la edad minima")
	}

}