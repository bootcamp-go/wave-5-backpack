package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(productos []Producto, res chan float64) {
	suma := .0
	for _, producto := range productos {
		suma += float64(producto.cantidad) * producto.precio
	}
	res <- suma
}

func sumarServicios(servicios []Servicio, res chan float64) {
	suma := .0
	for _, servicio := range servicios {
		if servicio.minutosTrabajados < 30 {
			servicio.minutosTrabajados = 30
		}
		suma += servicio.precio * float64(servicio.minutosTrabajados/30)
	}
	res <- suma
}

func sumarMantenimientos(mantenimientos []Mantenimiento, res chan float64) {
	suma := .0
	for _, mantenimiento := range mantenimientos {
		suma += mantenimiento.precio
	}
	res <- suma
}

func main() {
	productos := []Producto{
		{"Mouse", 12335.10, 2},
		{"Teclado", 123.5, 5},
	}

	servicios := []Servicio{
		{"Limpieza", 100, 68},
		{"Cocina", 254, 2},
	}

	mantenimientos := []Mantenimiento{
		{"PC", 56},
		{"Servidor", 135},
	}

	res := make(chan float64)

	go sumarProductos(productos, res)
	go sumarServicios(servicios, res)
	go sumarMantenimientos(mantenimientos, res)

	total := <-res + <-res + <-res

	fmt.Printf("Monto Total: $%.2f\n", total)
}
