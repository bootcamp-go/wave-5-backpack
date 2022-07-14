package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo         int64
	DNI            int64
	NombreApellido string
	Telefono       string
	Domcilio       string
}

func main() {
	var (
		dni            int64  = 0
		nombreapellido string = "Juan"
		telefono       string = "1239021"
		domicilio      string = "Valle"
	)

	legajoId, err := generarId()
	if err != nil {
		panic(err)
	}

	verificarCliente(legajoId)
	NewCliente(legajoId, dni, nombreapellido, telefono, domicilio)

	fmt.Println("Fin de la ejecución")
}

func generarId() (int64, error) {
	var id int = rand.Int()
	return int64(id), nil
}

func verificarCliente(id int64) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("No han quedado archivos abiertos")
	}()

	read, err := os.Open("./customer.txt")
	if err != nil {
		fmt.Println("Llegando al panic . . . 👀")
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	defer read.Close()
}

func NewCliente(legajo, dni int64, nombreapellido, telefono, domcilioo string) (*Cliente, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
	}()

	if legajo != 0 {
		panic("Legajo no puede ser 0")
	}

	if dni != 0 {
		panic("DNI no puede ser 0")
	}

	if nombreapellido != "" {
		panic("Nombre y apellido son requeridos")
	}

	if telefono != "" {
		panic("Telefono son requeridos")
	}

	if domcilioo != "" {
		panic("Domicilio son requeridos")
	}

	return &Cliente{Legajo: legajo, DNI: dni, NombreApellido: nombreapellido, Telefono: telefono, Domcilio: domcilioo}, nil
}
