package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo          int
	NombreYapellido string
	Dni             string
	NumeroTelefono  string
	Domicilio       string
}

func generarLegajo() (int, error) {
	return rand.Intn(10), nil
}

func inicializarCliente() Cliente {
	fmt.Println("Función inicializarCliente()")

	var legajo, err = generarLegajo()
	if err != nil {
		panic("Error al generar el ID")
	}

	return Cliente{
		Legajo: legajo,
	}
}

func registrarCliente(cliente Cliente) {
	fmt.Println("Función registrarCliente()")
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := leerArchivo("customers.txt")

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}

}

func validarCliente(cliente Cliente) (bool, error) {
	fmt.Println("Función validarCliente()")
	if cliente.Dni == "" || cliente.Domicilio == "" || cliente.Legajo == 0 || cliente.NombreYapellido == "" || cliente.NumeroTelefono == "" {
		return false, errors.New("Alguno de los datos del cliente no fue cargado")
	}
	return true, nil
}

func main() {
	defer func() {
		fmt.Println("No han quedado archivos abiertos")
	}()

	defer func() {
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}()

	defer func() {
		fmt.Println("Fin de la ejecución")
	}()

	var cliente = inicializarCliente()

	registrarCliente(cliente)

	validarCliente(cliente)

	fmt.Println("Fin de ejecución main")
}

func leerArchivo(path string) (*os.File, error) {
	fileDescriptor, err := os.Open(path)
	return fileDescriptor, err
}
