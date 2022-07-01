package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generarID() int {
	/* return rand.Intn(99999) */
	return 0
}

func getCustomers() map[int]Cliente {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	data, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado.")
	}
	rows := strings.Split(string(data), "\n")
	clientes := make(map[int]Cliente)

	for _, row := range rows[1:] {
		values := strings.Split(row, ";")
		cliente := Cliente{}
		cliente.Legajo, _ = strconv.Atoi(values[0])
		cliente.Nombre = values[1]
		cliente.Apellido = values[2]
		cliente.DNI, _ = strconv.Atoi(values[3])
		cliente.Telefono, _ = strconv.Atoi(values[4])
		cliente.Domicilio = values[5]
		clientes[cliente.Legajo] = cliente
	}

	return clientes
}

func verificarExistencia(legajo int) bool {
	clientes := getCustomers()
	_, ok := clientes[legajo]
	return ok
}

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Telefono  int
	Domicilio string
}

func validarDatos(cliente Cliente) (bool, error) {
	if cliente.Legajo == 0 ||
		cliente.Nombre == "" ||
		cliente.Apellido == "" ||
		cliente.DNI == 0 ||
		cliente.Telefono == 0 ||
		cliente.Domicilio == "" {
		return false, errors.New("los datos no pueden ser nulos")
	}
	return true, nil
}

func main() {

	defer func() {
		err := recover()

		fmt.Println("Fin de la ejecución")
		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("No han quedado archivos abiertos")
	}()

	id := generarID()

	if id == 0 {
		panic("Error al generar id")
	}

	if verificarExistencia(id) {
		panic("El cliente ya existe")
	}

	cliente := Cliente{
		Legajo:    id,
		Nombre:    "Lionel",
		Apellido:  "Messi",
		DNI:       12345678,
		Telefono:  1154321234,
		Domicilio: "Rosario",
	}

	_, err := validarDatos(cliente)
	if err != nil {
		fmt.Println(err)
	}

}
