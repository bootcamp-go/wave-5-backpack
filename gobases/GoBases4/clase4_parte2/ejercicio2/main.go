package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo         int64
	NombreApellido string
	DNI            int64
	Telefono       string
	Domicilio      string
}

func generarId() (int64, error) {
	var id int64 = int64(rand.Int())
	return id, nil
}

func verificarCliente(id int64) {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	read, err := os.Open("../customers.txt")

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta danado")
	}
	defer read.Close()
}

func NewCliente(legajo, dni int64, nombreapellido, telefono, domicilio string) (*Cliente, error) {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
		}
	}()

	if legajo != 0 {
		panic("Legajo no puede ser 0")
	}
	if dni != 0 {
		panic("DNI no puede ser 0")
	}
	if nombreapellido != "" {
		panic("Nombre y apellidos son requeridos")
	}
	if telefono != "" {
		panic("Telefono requerido")
	}
	if domicilio != "" {
		panic("Direccion son requeridas")
	}

	return &Cliente{Legajo: legajo, DNI: dni, NombreApellido: nombreapellido, Telefono: telefono, Domicilio: domicilio}, nil
}

func main() {

	var (
		dni            int64  = 12345
		nombreapellido string = "Juan David Serna"
		telefono       string = "312424124"
		domicilio      string = "Tulua"
	)

	legajoId, err := generarId()

	if err != nil {
		panic(err)
	}

	verificarCliente(legajoId)
	NewCliente(legajoId, dni, nombreapellido, telefono, domicilio)

}
