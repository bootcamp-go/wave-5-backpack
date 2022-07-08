package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicios struct {
	Nombre            string
	Precio            int
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func SumarProductos(productos *[]Producto, c chan float64) {
	var total float64 = 0
	for _, producto := range *productos {
		total = total + (float64(producto.Precio) + float64(producto.Cantidad))
	}

	fmt.Println("Total productos $", total)

	c <- total
}

func SumarServicios(servicios *[]Servicios, c chan float64) {
	var total float64 = 0

	for _, servicio := range *servicios {
		var tiempo float64 = float64(servicio.MinutosTrabajados) / 30
		if servicio.MinutosTrabajados < 30 {
			tiempo = 1
		}
		total += (float64(servicio.Precio) * tiempo)
	}
	fmt.Println("Total servicios $", total)

	c <- total

}

func SumarMantenimiento(mantenimientos *[]Mantenimiento, c chan float64) {
	var total float64 = 0
	for _, value := range *mantenimientos {
		total += float64(value.Precio)
	}

	fmt.Println("Total mantenimiento $", total)
	c <- total
}

func main() {
	productos := []Producto{
		{Nombre: "Producto 1", Precio: 100, Cantidad: 20},
		{Nombre: "Producto 2", Precio: 200, Cantidad: 4},
		{Nombre: "Producto 3", Precio: 300, Cantidad: 1},
		{Nombre: "Producto 4", Precio: 200, Cantidad: 200},
	}

	servicios := []Servicios{
		{Nombre: "Programación", Precio: 200, MinutosTrabajados: 480},
		{Nombre: "Paqueteria", Precio: 10234, MinutosTrabajados: 30},
		{Nombre: "Encomienda", Precio: 300, MinutosTrabajados: 15},
		{Nombre: "Limpieza", Precio: 100, MinutosTrabajados: 22},
	}

	mantenimientos := []Mantenimiento{
		{Nombre: "Mantenimiento 1", Precio: 400},
		{Nombre: "Mantenimiento 2", Precio: 2500},
		{Nombre: "Mantenimiento 3", Precio: 4100},
		{Nombre: "Mantenimiento 4", Precio: 4100},
	}

	// Usando tres canales
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go SumarProductos(&productos, c1)
	go SumarServicios(&servicios, c2)
	go SumarMantenimiento(&mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("* Total final * $ ", t1+t2+t3)
}
