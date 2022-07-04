package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase3/tt/ejercicio3/calculo"
)

func main() {
  productos := []calculo.Productos{
    {Nombre: "mate", Precio: 150.5, Cantidad: 5},
    {Nombre: "termo", Precio: 300, Cantidad: 1},
    {Nombre: "bombilla", Precio: 45.5, Cantidad: 2},
  }

  servicios := []calculo.Servicios{
    {Nombre: "lavado", Precio: 100, MinutosTrabajados: 31},
    {Nombre: "lavado", Precio: 100, MinutosTrabajados: 90},
    {Nombre: "lavado", Precio: 100, MinutosTrabajados: 91},
  }

  mantenimiento := []calculo.Mantenimiento{
    {Nombre: "basico", Precio:100},
    {Nombre: "extra", Precio: 125.5},
    {Nombre: "premium", Precio: 190.3},
  }

  p := make(chan float64)
  go calculo.SumarProductos(productos, p)

  s := make(chan float64)
  go calculo.SumarServicios(servicios, s)

  m := make(chan float64)
  go calculo.SumarMantenimiento(mantenimiento, m)

	total := <-p + <-s + <-m

	fmt.Printf("Total: %v\n", total)
}
