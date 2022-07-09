package main

import "fmt"

type Client struct {
	edad           int
	empleadoStatus bool
	antiguedad     float32
	sueldo         int
}

func main() {
	newClient := Client{
		edad:           23,
		empleadoStatus: true,
		antiguedad:     2,
		sueldo:         9000,
	}
	bankProcess(newClient)
}

func bankProcess(cliente Client) {
	if cliente.edad < 22 {
		fmt.Println("Debes ser mayor de edad para solicitar un préstamo")
		return
	}
	if cliente.empleadoStatus != true {
		fmt.Println("Debes ser empleado actualmente para solicitar un préstamo")
		return
	}
	if cliente.antiguedad < 1 {
		fmt.Println("Debes tener más de 1 año de antiguedad para solicitar un préstamo")
		return
	}
	if cliente.sueldo < 10000 {
		fmt.Println("Interés 0 en tu préstamo")
	}

	fmt.Println("Cumples con las condiciones para el prestamo")

}
