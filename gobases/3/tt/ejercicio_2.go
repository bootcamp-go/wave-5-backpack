package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Usuario struct {
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Correo    string     `json:"correo"`
	Productos []Producto `json:"productos"`
}

type Producto struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {
	// Cargar un usuario
	u := Usuario{
		Nombre:    "Agustin",
		Apellido:  "Nigrelli",
		Correo:    "agustin.nigrelli@mercadolibre.com",
		Productos: nil,
	}

	sin_productos, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nUsuario sin productos\n" + string(sin_productos))

	// Crear productos y agregarlos al usuario
	p1 := NuevoProducto("Jabon", 5.50)
	p2 := NuevoProducto("Detergente", 10)
	p3 := NuevoProducto("Lavandina", 7.25)

	Agregar(&u, p1, 3)
	Agregar(&u, p2, 5)
	Agregar(&u, p3, 10)

	con_productos, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nUsuario con productos\n" + string(con_productos))

	// Borrar los productos del usuario
	Borrar(&u)
	productos_borrados, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nUsuario con productos borrados\n" + string(productos_borrados))

}

func NuevoProducto(n string, p float64) Producto {
	return Producto{Nombre: n, Precio: p}
}

func Agregar(u *Usuario, p Producto, c int) {
	p.Cantidad = c
	u.Productos = append(u.Productos, p)
}

func Borrar(u *Usuario) {
	u.Productos = nil
}
