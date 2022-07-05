package main

import "fmt"

var cliente struct {
	edad         int
	empleado     bool
	tiempoEmpleo int
	salario      float32
}

func main() {

	if cliente.edad > 22 && cliente.empleado && cliente.tiempoEmpleo > 1 {
		if cliente.salario > 100000 {
			fmt.Println("Tu prestamo ha sido aprobado y tendra 0 de interes")
		}
		fmt.Println("Tu prestamo ha sido aprobado")
	}
	fmt.Println("No cumples con los requisitos para el prestamo")

}
