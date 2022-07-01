package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cliente struct {
	Legajo    string
	Nombre    string
	DNI       string
	telefono  string
	domicilio string
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func generarLegajo() string {
	if RandBool() {
		var vacio string
		return vacio
	}
	return "dja4324d-dsajkddw23"
}

func readFile(name string) *os.File {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("No han quedado archivos abiertos")
	}()

	file, err := os.OpenFile(name, os.O_RDONLY, 0600)

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	return file
}

func registrarCliente(legajo string, nombre string, dni string, telefono string, domicilio string) (Cliente, error) {

	if legajo == "" {
		err := fmt.Errorf("error: se ingreso legajo vacio")
		return Cliente{}, err
	}

	if nombre == "" {
		err := fmt.Errorf("error: se ingreso nombre vacio")
		return Cliente{}, err
	}

	if dni == "" {
		err := fmt.Errorf("error: se ingreso dni vacio")
		return Cliente{}, err
	}

	if len(telefono) != 9 {
		err := fmt.Errorf("error: telefono debe contener 9 numeros: %s", telefono)
		return Cliente{}, err
	}

	if domicilio <= "" {
		err := fmt.Errorf("error: se ingreso domicilio vacio")
		return Cliente{}, err
	}

	return Cliente{Legajo: legajo, Nombre: nombre, DNI: dni, telefono: telefono, domicilio: domicilio}, nil
}

func main() {

	defer func() {
		err := recover()
		fmt.Println("Fin de la ejecucion")
		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("No han quedado archivos abiertos")
	}()
	legajo := generarLegajo()
	if legajo == "" {
		panic("error legajo esta vacio")
	}

	cliente, err := registrarCliente(legajo, "Claudio Figueroa", "19572311-7", "321312", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cliente Creado exitosamente", cliente)
	}

}
