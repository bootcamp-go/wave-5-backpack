package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       string
	Telefono  string
	Domicilio string
}

func newCliente(legajo int, nombre string, apellido string, dNI string, telefono string, domicilio string) cliente {
	return cliente{
		Legajo:    legajo,
		Nombre:    nombre,
		Apellido:  apellido,
		DNI:       dNI,
		Telefono:  telefono,
		Domicilio: domicilio,
	}
}

func generarNumeroLegajo() (int, error) {
	id := rand.Int()

	if id <= 0 {
		return 0, errors.New("ID invalido")
	} else {
		return id, nil
	}
}

func verificarCliente(legajo int) bool {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile("customers.txt")

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		return true
	}
}

func controlDatos(cliente cliente) (cliente, error) {
	if cliente.Legajo == 0 || cliente.Nombre == "" || cliente.Apellido == "" || cliente.DNI == "" ||
		cliente.Telefono == "" || cliente.Domicilio == "" {
		return cliente, errors.New("error: todos los datos deben contener un valor")
	} else {
		return cliente, nil
	}
}
