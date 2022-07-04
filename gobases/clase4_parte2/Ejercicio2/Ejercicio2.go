package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	Nombre, Apellido, Telefono, Domicilio string
	Legajo                                int
}

func leerArchivo(path string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := os.ReadFile(path)
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	fmt.Println("Archivo leido con exito")
}

func crearIdLegajo(i *int) int {
	result := *i
	*i++
	return result
}

func verificarCliente(cliente *Cliente) (bool, error) {
	if cliente.Nombre == "" || cliente.Apellido == "" || cliente.Telefono == "" || cliente.Domicilio == "" || cliente.Legajo == 0 {
		return false, errors.New("Algun valor no esta inicializado")
	}
	return true, nil
}

func main() {
	var i int = 1
	a := crearIdLegajo(&i)
	b := crearIdLegajo(&i)
	c := crearIdLegajo(&i)
	fmt.Println(a, b, c, i)
}
