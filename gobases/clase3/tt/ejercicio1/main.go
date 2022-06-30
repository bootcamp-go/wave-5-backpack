package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase3/tt/ejercicio1/usuarios"
)

func main() {
  usuario := &usuarios.Usuarios{
    Nombre: "Magnus",
    Apellido: "Carlsen",
  }

  fmt.Printf("Datos del usuario: %+v\n", usuario)

  usuario.CambiarNombre("Pepe","Cuenca")
  usuario.CambiarEdad(35)
  usuario.CambiarCorreo("pepe@email.com")
  usuario.CambiarPassword("12345")

  fmt.Printf("Datos nuevo del usuario: %+v\n", usuario)
}
