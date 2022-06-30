package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase2/tt/ejercicio3/condiciones"
)

func main() {
	mate, err := condiciones.NuevoProducto("pequeno", "mate", 100)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", mate)
}
