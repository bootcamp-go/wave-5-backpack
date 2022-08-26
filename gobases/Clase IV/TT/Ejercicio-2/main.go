package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliente struct {
	Legajo          int
	NombreYApellido string
	DNI             int
	NumeroTelefono  int
	Domicilio       string
}

var (
	Name      string = "Pedro Gonzalez"
	DNI       int    = 0
	Telefono  int    = 971596304
	Domicilio string = "Perez Rosales 1384,Las Condes,RM"
)

func main() {

	legajoID, err := generateIdLegajo()
	if err != nil {
		panic("error en la generación del legajo")
	}

	verificarClienteExiste(legajoID)
	err, _ = validarDatosCliente(Name, Domicilio, legajoID, DNI, Telefono)
	if err != nil {
		panic("error en la generación del cliente")
	}

	fmt.Println("Fin de la ejecucición")

}

func generateIdLegajo() (int, error) {
	id := rand.Intn(999999999)
	return id, nil
}

func verificarClienteExiste(legajoID int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("No han quedado archivos abiertos")

	}()
	_, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func validarDatosCliente(name, domicilio string, legajo, dni, telefono int) (error, *cliente) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Se detectaron errores en tiempo de ejecución")
		}
	}()
	if name == "" {
		panic("nombre ingresado está vacío")

	}
	if domicilio == "" {
		panic("domicilio ingresado está vacío")

	}
	if legajo == 0 {
		panic(" dni ingresado está vacío")

	}
	if dni == 0 {
		panic("dni ingresado está vacío")

	}
	if telefono == 0 {
		panic("teléfono ingresado está vacío")

	}

	return nil, &cliente{Legajo: legajo, NombreYApellido: name, Domicilio: domicilio, DNI: dni, NumeroTelefono: telefono}
}
