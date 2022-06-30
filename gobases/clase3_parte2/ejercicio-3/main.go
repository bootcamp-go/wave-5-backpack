package main

import "fmt"

// ===========================================
// ================= Structs =================
// ===========================================

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

// ===========================================
// ================ Funciones ================
// ===========================================

func SumarProductos(productos []Productos, c chan float64) {
	var total float64
	for _, producto := range productos {
		total += producto.precio * float64(producto.cantidad)
	}
	c <- total
}

func SumarServicios(servicios []Servicios, c chan float64) {
	var total float64
	for _, servicio := range servicios {
		if servicio.minutosTrabajados < 30 {
			total += servicio.precio * float64(30)
		} else {
			total += servicio.precio * float64(servicio.minutosTrabajados)
		}
	}
	c <- total
}

func SumarMantenimientos(mantenimientos []Mantenimiento, c chan float64) {
	var total float64
	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.precio
	}
	c <- total
}

// ===========================================
// ================== Main ===================
// ===========================================

func main() {

	var productos []Productos = []Productos{
		{nombre: "Laptop", precio: 1000, cantidad: 1},
		{nombre: "Mouse", precio: 50, cantidad: 2},
		{nombre: "Teclado", precio: 100, cantidad: 3},
	}

	var servicios []Servicios = []Servicios{
		{nombre: "Limpieza", precio: 100, minutosTrabajados: 30},
		{nombre: "Reparacion", precio: 200, minutosTrabajados: 60},
		{nombre: "Mantenimiento", precio: 300, minutosTrabajados: 90},
	}

	var mantenimientos []Mantenimiento = []Mantenimiento{
		{nombre: "Limpieza", precio: 100},
		{nombre: "Reparacion", precio: 200},
		{nombre: "Mantenimiento", precio: 300},
	}

	c := make(chan float64)

	go SumarProductos(productos, c)
	go SumarServicios(servicios, c)
	go SumarMantenimientos(mantenimientos, c)

	total := <-c + <-c + <-c

	fmt.Println("Total:", total)
}
