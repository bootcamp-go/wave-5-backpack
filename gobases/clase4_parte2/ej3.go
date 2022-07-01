package main

import (
	"fmt"
	"math"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func main() {
	productos := []Producto{
		{"Caramelo", 5.0, 10},
		{"Chupetin", 15.0, 2},
		{"Alfajor", 50.0, 1},
	}

	servicios := []Servicio{
		{"Limpieza", 100, 150},
		{"Desinfección", 50, 25},
		{"Salud", 500, 90},
	}

	mantenimientos := []Mantenimiento{
		{"Reparaciones varias", 500},
		{"Aires acondicionados", 3000},
	}

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go sumarProductos(productos, c1)
	go sumarServicios(servicios, c2)
	go sumarMantenimientos(mantenimientos, c3)

	fmt.Println("Productos: ", <-c1)
	fmt.Println("Servicios: ", <-c2)
	fmt.Println("Mantenimientos: ", <-c3)

}

func sumarProductos(productos []Producto, c chan float64) {
	var total float64 = 0

	for _, producto := range productos {
		total += producto.Precio * float64(producto.Cantidad)
	}
	c <- total
}

func sumarServicios(servicios []Servicio, c chan float64) {
	var total float64 = 0

	for _, servicio := range servicios {
		var mediasHorasTrabajadas float64 = math.Max(1, servicio.MinutosTrabajados/30)
		total += servicio.Precio * mediasHorasTrabajadas
	}
	c <- total
}

func sumarMantenimientos(mantenimientos []Mantenimiento, c chan float64) {
	var total float64 = 0

	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.Precio
	}
	c <- total
}
