package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type Servicios struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados uint64
}
type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos *[]Producto, c chan float64) {

	var total float64
	for _, value := range *productos {
		total += value.Precio * float64(value.Cantidad)
	}

	fmt.Println("El total de Productos es $", total)
	c <- total
	close(c)
}

func sumarServicios(servicio *[]Servicios, c chan float64) {

	var total float64
	for _, value := range *servicio {
		if value.MinutosTrabajados < 30 {
			total += value.Precio * 30
		} else {
			total += value.Precio * float64(value.MinutosTrabajados)
		}
	}
	fmt.Println("EL total de servicios es $", total)
	c <- total
	close(c)
}

func sumarMantenimiento(mantenimiento *[]Mantenimiento, c chan float64) {
	var total float64
	for _, value := range *mantenimiento {
		total += value.Precio
	}

	fmt.Println("El total de mantenimientos es $", total)
	c <- total
	close(c)
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

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go sumarProductos(&productos, c1)
	go sumarServicios(&servicios, c2)
	go sumarMantenimiento(&mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("El total final es $", t1+t2+t3)

}
