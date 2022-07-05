package main

import "fmt"

type Productos struct{
	Nombre string
	Precio int
	Cantidad int
}
type Servicios struct{
	Nombre string
	Precio int
	Minutos int
}
type Mantenimiento struct{
	Nombre string
	Precio int
}

func sumarProductos(productos []Productos, c chan int){
	suma := 0
	for _, value := range productos {
		suma += value.Precio * value.Cantidad
	}
	c <- suma
	close(c)
}
func sumarServicios(servicios []Servicios, c chan int){
	suma := 0
	for _, value := range servicios {
		suma += value.Precio * value.Minutos
	}
	c <- suma
	close(c)
}
func sumarMantenimientos(mantenimiento []Mantenimiento, c chan int){
	suma := 0
	for _, value := range mantenimiento {
		suma += value.Precio 
	}
	c <- suma
	close(c)
}

func main(){
	productos := []Productos{
		{Nombre: "Producto 1", Precio: 100, Cantidad: 20},
		{Nombre: "Producto 2", Precio: 200, Cantidad: 4},
		{Nombre: "Producto 3", Precio: 300, Cantidad: 1},
		{Nombre: "Producto 4", Precio: 200, Cantidad: 200},
	}

	servicios := []Servicios{
		{Nombre: "Programación", Precio: 200, Minutos: 480},
		{Nombre: "Paqueteria", Precio: 10234, Minutos: 30},
		{Nombre: "Encomienda", Precio: 300, Minutos: 15},
		{Nombre: "Limpieza", Precio: 100, Minutos: 22},
	}

	mantenimientos := []Mantenimiento{
		{Nombre: "Mantenimiento 1", Precio: 400},
		{Nombre: "Mantenimiento 2", Precio: 2500},
		{Nombre: "Mantenimiento 3", Precio: 4100},
		{Nombre: "Mantenimiento 4", Precio: 4100},
	}

	// Usando tres canales
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go sumarProductos(productos, c1)
	go sumarServicios(servicios, c2)
	go sumarMantenimientos(mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("* Total final * $ ", t1+t2+t3)

}