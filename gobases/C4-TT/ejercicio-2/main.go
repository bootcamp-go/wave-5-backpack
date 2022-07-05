package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo          int
	DNI             int
	Nombre_Apellido string
	Telefono        string
	Domcilio        string
}

func main() {
	var (
		dni             int    = 42263580
		nombre_apellido string = "Matias Vince"
		telefono        string = "3435508581"
		domicilio       string = "Estacion sosa"
	)

	legajo, err := generarLegajo()
	if err != nil {
		panic(err)
	}

	verificarCliente(legajo)
	nuevoCliente(legajo, dni, nombre_apellido, telefono, domicilio)

	fmt.Println("Fin de la ejecucion")
}

func generarLegajo() (int, error) {
	var legajo int = rand.Int()
	return legajo, nil
}

func verificarCliente(legajo int) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile("dir")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func nuevoCliente(legajo, dni int, nombre_apellido, telefono, domicilio string) (Cliente, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	if legajo == 0 {
		panic("Legajo no puede ser 0")
	}

	if dni == 0 {
		panic("DNI no puede ser 0")
	}

	if nombre_apellido == "" {
		panic("Debe ingresar un nombre y apellido")
	}

	if telefono == "" {
		panic("Debe ingresar un telefono")
	}

	if domicilio == "" {
		panic("Debe ingresar un domicilio")
	}

	return Cliente{
		Legajo:          legajo,
		DNI:             dni,
		Nombre_Apellido: nombre_apellido,
		Telefono:        telefono,
		Domcilio:        domicilio,
	}, nil
}
