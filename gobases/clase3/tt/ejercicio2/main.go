package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase3/tt/ejercicio2/ecommerce"
)

func main() {
	usuario := ecommerce.NuevoUsuario("Pepe","Cuenca","pepe@email.com")
	fmt.Printf("Iniciando usuario %+v\n", usuario)	

	// Crea producto mate
	mate := ecommerce.NuevoProducto("mate", 100)

	// Agrega 1 mates al usuario
	ecommerce.AgregarProducto(&usuario, &mate, 1)

	for _, v := range *usuario.Productos {
		fmt.Printf("Producto: %+v\n", v)
	}

	// Agrega 2 mates al usuario
	ecommerce.AgregarProducto(&usuario, &mate, 2)

	for _, v := range *usuario.Productos {
		fmt.Printf("Producto: %+v\n", v)
	}
	
	ecommerce.BorrarProductos(&usuario)
	fmt.Println(*usuario.Productos)
}
