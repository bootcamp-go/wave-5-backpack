package main

import (
	"math"
)

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            int
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func SumarProductos(productos *[]*Producto, resultado chan int) {
	suma := 0
	for _, producto := range *productos {
		suma += producto.Precio * producto.Cantidad
	}
	resultado <- suma
}

func SumarServicios(servicios *[]*Servicio, resultado chan int) {
	suma := 0
	for _, servicios := range *servicios {
		cantidadDePagos := int(math.Floor(float64(servicios.MinutosTrabajados / 30)))
		resto := servicios.MinutosTrabajados % 30
		if resto > 0 {
			cantidadDePagos += 1
		}
		suma += servicios.Precio * cantidadDePagos
	}
	resultado <- suma
}

func SumarMantenimiento(mantenimientos *[]*Mantenimiento, resultado chan int) {
	suma := 0
	for _, mantenimiento := range *mantenimientos {
		suma += mantenimiento.Precio
	}
	resultado <- suma
}

func main() {
	sumaProducto, sumaServicio, sumaMantenimiento := make(chan int), make(chan int), make(chan int)
	productos := []*Producto{
		{Nombre: "Iphone s5", Precio: 12000, Cantidad: 10},
		{Nombre: "Pantalla 55'", Precio: 200000, Cantidad: 1},
		{Nombre: "Pantalla 27'", Precio: 120000, Cantidad: 3},
	}
	servicios := []*Servicio{
		{Nombre: "Limpieza", Precio: 10000, MinutosTrabajados: 80},
		{Nombre: "Mantenimiento", Precio: 5000, MinutosTrabajados: 120},
		{Nombre: "Gestion de video", Precio: 20000, MinutosTrabajados: 30},
	}
	mantenimientos := []*Mantenimiento{
		{Nombre: "Maquinas", Precio: 10000},
		{Nombre: "Radios", Precio: 3500},
		{Nombre: "Comida", Precio: 5000},
	}

	go SumarProductos(&productos, sumaProducto)
	go SumarMantenimiento(&mantenimientos, sumaMantenimiento)
	go SumarServicios(&servicios, sumaServicio)

	resulProd, resulMant, resulServc := <-sumaProducto, <-sumaMantenimiento, <-sumaServicio
	println("Total Productos:    ", resulProd)
	println("Total Servicios:    ", resulMant)
	println("Total Mantenimiento:", resulServc)

	println("Total:              ", resulProd+resulMant+resulServc)
}
