package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	legajo    int64
	nombre    string
	apellido  string
	DNI       int
	telefono  string
	domicilio string
}

type ClienteList struct {
	client []Cliente
}

func addCliente(legajo int64, nombre string, apellido string, DNI int, telefono, domcilio string) (*Cliente, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
	}()
	if legajo != 0 {
		panic("El legajo no puede ser cero")
	}
	if nombre != "" {
		panic("El nombre no puede ser vacio")
	}
	if apellido != "" {
		panic("El apellido no puede ser vacio")
	}
	if DNI != 0 {
		panic("El DNI no puede ser 0")
	}
	if telefono != "" {
		panic("El telefono no puede ser 0")
	}
	if domcilio != "" {
		panic("La direccion no puede ser vacia")
	}
	return &Cliente{legajo: legajo, nombre: nombre, apellido: apellido, DNI: DNI, telefono: telefono, domicilio: domcilio}, nil
}

func listClients(c *ClienteList, clie Cliente) {
	c.client = append(c.client, clie)
}

func generadorLegajo() (int64, error) {
	var legajo int = rand.Int()
	return int64(legajo), nil
}

func verificarCliente(id int64) {
	defer func() {
		err := recover() // Recupera el panic - para evitar una ejecución no deseada
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("No han quedado archivos abiertos")
	}()
	read, err := os.Open("./customer.txt")
	if err != nil {
		fmt.Println("Llegando al panic . . . 👀")
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	defer read.Close()
	// Si existe el archivo validamos - leyendo el archivo - (Opcional)
}

func main() {

	var (
		nombre    string = "Eminem"
		apellido  string = "Anuel"
		DNI       int    = 1005
		telefono  string = "3008529053"
		domicilio string = "Cajamarca"
	)
	legajo, err := generadorLegajo()
	if err != nil {
		panic(err)
	}

	verificarCliente(legajo)
	addCliente(legajo, nombre, apellido, DNI, telefono, domicilio)
	fmt.Println("Fin de la ejecución")
}
