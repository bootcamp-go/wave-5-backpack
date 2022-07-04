package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Cliente struct {
	Legajo    string
	Nombre    string
	Apellido  string
	DNI       int
	Telefono  int
	Domicilio string
}

func main() {
	legajo := IdGen()
	if legajo == "" {
		panic("IdGen() ha producido un id nulo")
	}

	nombre := "Agustin"
	apellido := "Nigrelli"
	dni := 0
	telefono := 1155555555
	domicilio := "Polo DOT 123"

	ValidarLegajo(legajo)

	// Pongo la variable en _ para que no salga error de variable sin utilizar
	_, err := GenerarCliente(legajo, nombre, apellido, dni, telefono, domicilio)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fin de la ejecución")
}

func IdGen() string {
	length := 10
	rand.Seed(time.Now().Unix())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func ValidarLegajo(legajo string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("No han quedado archivos abiertos")
	}()

	file, err := os.Open("./gobases/4/tt/customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	defer file.Close()

	reader, err := io.ReadAll(file)

	lines := strings.Split(string(reader), "\n")
	cols := strings.Split(lines[1], ",")

	if legajo == cols[0] || legajo == "" {
		panic("el numero de legajo ya existe o es invalido")
	}
}

func GenerarCliente(legajo string, nombre string, apellido string, dni int, telefono int, domicilio string) (c *Cliente, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}()

	if nombre == "" || apellido == "" || dni == 0 || telefono == 0 || domicilio == "" {
		panic("alguno de los campos ingresados es nulo")
	}

	c = &Cliente{
		Legajo:    legajo,
		Nombre:    nombre,
		Apellido:  apellido,
		DNI:       dni,
		Telefono:  telefono,
		Domicilio: domicilio,
	}

	return c, nil
}
