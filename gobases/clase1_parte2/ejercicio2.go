package main

import "fmt"

type Cliente struct {
	nombre           string
	edad             int
	empleado         bool
	antiguedad_meses int
	sueldo           int
}

func main() {
	cliente1 := Cliente{"Juan", 20, true, 11, 50000}
	cliente2 := Cliente{"Pedro", 23, false, 0, 10000}
	cliente3 := Cliente{"Pablo", 40, true, 7, 45000}
	cliente4 := Cliente{"Lucas", 30, true, 24, 80000}
	cliente5 := Cliente{"Marcos", 60, true, 36, 110000}
	var clientes = []Cliente{cliente1, cliente2, cliente3, cliente4, cliente5}

	for idx := range clientes {
		cliente := clientes[idx]
		checkCredito(cliente)
	}
}

func checkCredito(cliente Cliente) {
	if cliente.edad > 22 {
		if cliente.empleado {
			if cliente.antiguedad_meses > 12 {
				fmt.Printf("El cliente %s es apto para un crédito!\n", cliente.nombre)
				if cliente.sueldo > 100000 {
					fmt.Printf("Felicidades, el cliente %s no debe pagar intereses\n", cliente.nombre)
				}
			} else {
				fmt.Printf("El cliente %s no tiene mas de un año de antiguedad en su trabajo\n", cliente.nombre)
			}
		} else {
			fmt.Printf("El cliente %s no tiene empleo\n", cliente.nombre)
		}
	} else {
		fmt.Printf("El cliente %s no tiene la edad suficiente \n", cliente.nombre)
	}
}
