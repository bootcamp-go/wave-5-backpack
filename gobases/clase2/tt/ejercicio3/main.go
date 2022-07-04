package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase2/tt/ejercicio3/condiciones"
)

func main() {
	bombilla, err := condiciones.NuevoProducto("pequeno", "bombilla", 100) // total = precio
	if err != nil {
		log.Fatal(err)
	}

	mate, err := condiciones.NuevoProducto("mediano", "mate", 300) // total = precio * 0.03
	if err != nil {
		log.Fatal(err)
	}

	mochilaViajera, err := condiciones.NuevoProducto("grande", "mochila viajera", 1000) // total = 2500 + precio * 0.06

	tienda := condiciones.NuevaTienda()
	fmt.Printf("Total en tienda: %v\n", tienda.Total())

	tienda.Agregar(bombilla)
	fmt.Printf("Total en tienda: %v\n", tienda.Total())

	tienda.Agregar(mate)
	fmt.Printf("Total en tienda: %v\n", tienda.Total())

	tienda.Agregar(mochilaViajera)
	fmt.Printf("Total en tienda: %v\n", tienda.Total())
}
