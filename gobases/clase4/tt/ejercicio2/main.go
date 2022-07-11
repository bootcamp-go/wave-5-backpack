package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase4/tt/ejercicio2/clientes"
)

func main() {
	cliente := &clientes.Cliente{
		Legajo:    1,
		Nombre:    "Pepe",
		Apellido:  "Cuenca",
		DNI:       12345678,
		Domicilio: "Wallaby 42 Sydney",
		NumeroTel: "1122334455",
	}

	id := clientes.NuevoId()
	if id == 0 {
		panic("id retornado es 0")
	}

	err := clientes.CheckId(id)
	if err != nil {
		log.Println(err)
	} else {
		nuevoCliente, err := clientes.NuevoCliente(cliente.Legajo, cliente.DNI, cliente.Nombre, cliente.Apellido, cliente.Domicilio, cliente.NumeroTel)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("nuevo cliente: %+v\n", nuevoCliente)
		}
	}

	log.Println("Fin de la ejecución")
}
