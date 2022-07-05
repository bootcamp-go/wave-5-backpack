package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo         int
	DNI            int
	NombreApellido string
	Telefono       string
	Domcilio       string
}

func generarId() (int, error) {
	var id int = rand.Int()
	return id, nil
}

func verificarCliente(id int) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	read, err := os.Open("./customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	defer read.Close()
}

func NewCliente(legajo, dni int, nombreapellido, telefono, domcilioo string) (*Cliente, error) {
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

func main() {
	var (
		dni            int    = 1
		nombreapellido string = "Kevin Sossa"
		telefono       string = "322524"
		domicilio      string = "Cra 74"
	)

	legajoId, err := generarId()
	if err != nil {
		panic(err)
	}

	verificarCliente(legajoId)
	NewCliente(legajoId, dni, nombreapellido, telefono, domicilio)

	fmt.Println("Fin de la ejecución")
}
