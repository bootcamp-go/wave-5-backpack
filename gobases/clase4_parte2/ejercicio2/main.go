package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cliente struct {
	Legajo                                     int
	Nombre, Apellido, DNI, Telefono, Domicilio string
}

func asignarLegajo() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Int()
	//return rand.Intn(0) // GENERA PANIC
}

func validarClientePrevio(legajo int) (bool, error) {
	_, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañado")
	}
	return true, nil
}

func validarClientePrevioWrapper(legajo int) (bool, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return validarClientePrevio(legajo)
}
func validarDatosCliente(nombre, apellido, dni, telefono, domicilio string) (bool, error) {
	if nombre == "" || apellido == "" || dni == "" || telefono == "" || domicilio == "" {
		return false, errors.New("Datos invalidos, por favor revisar")
	}
	return true, nil
}

func agregarCliente(nombre, apellido, dni, telefono, domicilio string) (Cliente, error) {
	legajo := asignarLegajo()
	if legajo == 0 {
		panic("Error asignando legajo")
	}
	validarClientePrevioWrapper(legajo)
	_, err := validarDatosCliente(nombre, apellido, dni, telefono, domicilio)
	if err != nil {
		return Cliente{}, err
	}
	return (Cliente{legajo, nombre, apellido, dni, telefono, domicilio}), nil
}

func main() {
	defer func() {
		fmt.Println("No han quedado archivos abiertos")
	}()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecucion") // solo se imprime si falla la funcion asignarLegajo, la unica no controlada
		}
	}()
	defer func() {
		fmt.Println("Fin de la ejecucion")
	}()

	c, err := agregarCliente("Cristobal", "Monsalve", "123", "1234", "calle falsa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
