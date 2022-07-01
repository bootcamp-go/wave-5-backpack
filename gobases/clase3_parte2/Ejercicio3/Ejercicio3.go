package main

import "time"

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  int
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func sumarProductos(productos []Producto, c chan int) {
	var resultado int
	for _, producto := range productos {
		resultado += producto.Precio * producto.Cantidad
	}
	c <- resultado
}

func sumarServicios(servicios []Servicio, c chan int) {
	var resultado int
	for _, servicio := range servicios {
		mediasHoraTrabajadas := servicio.Minutos / 30
		if servicio.Minutos%30 > 0 {
			mediasHoraTrabajadas += 1
		}
		resultado += mediasHoraTrabajadas * (servicio.Precio / 2)
	}
	c <- resultado
}

func sumarMantenimiento(mantenimientos []Mantenimiento, c chan int) {
	var resultado int
	for _, mantenimiento := range mantenimientos {
		resultado += mantenimiento.Precio
	}
	c <- resultado
}

func main() {
	var productos = []Producto{
		{
			Nombre:   "a",
			Precio:   10,
			Cantidad: 5,
		},
		{
			Nombre:   "b",
			Precio:   7,
			Cantidad: 3,
		},
		{
			Nombre:   "c",
			Precio:   5,
			Cantidad: 3,
		},
	}
	var servicios = []Servicio{
		{
			Nombre:  "d",
			Precio:  10,
			Minutos: 30,
		},
		{
			Nombre:  "e",
			Precio:  8,
			Minutos: 30,
		},
		{
			Nombre:  "f",
			Precio:  6,
			Minutos: 30,
		},
	}
	var mantenimientos = []Mantenimiento{
		{
			Nombre: "g",
			Precio: 4,
		},
		{
			Nombre: "h",
			Precio: 6,
		},
		{
			Nombre: "i",
			Precio: 10,
		},
	}
	c := make(chan int)
	go sumarProductos(productos, c)
	go sumarServicios(servicios, c)
	go sumarMantenimiento(mantenimientos, c)
	time.Sleep(time.Second)
	var resultado int
	for i := 0; i < 3; i++ {
		resultado += <-c
	}
	println(resultado)
}
