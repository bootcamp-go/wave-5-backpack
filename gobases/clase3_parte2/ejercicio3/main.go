/*Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
 Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la
 sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
	 si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/
package main

import "fmt"

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)
	ArrayProductos := []Productos{
		{Nombre: "Play", Precio: 1000, cantidad: 2},
		{Nombre: "Xbox", Precio: 500, cantidad: 1},
	}

	ArrayMantenimientos := []Mantenimiento{
		{Nombre: "Arreglo", Precio: 10},
		{Nombre: "Limpieza", Precio: 20},
	}

	ArrayServicios := []Servicios{
		{Nombre: "ComprarJuegos", Precio: 50, MinsTrab: 1},
		{Nombre: "Comprar Plus", Precio: 10, MinsTrab: 1},
	}

	go SumarProductos(ArrayProductos, c1)
	t1 := <-c1
	go SumarMantenimiento(ArrayMantenimientos, c2)
	t2 := <-c2
	go SumarServicios(ArrayServicios, c3)
	t3 := <-c3
	fmt.Println("El total de los 3 es de: ", t1+t2+t3)
}

type Productos struct {
	Nombre   string
	Precio   float64
	cantidad int
}

type Servicios struct {
	Nombre   string
	Precio   float64
	MinsTrab float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(producs []Productos, c1 chan float64) {
	total := 0.0
	for _, prodi := range producs {
		total += prodi.Precio * float64(prodi.cantidad)
	}
	c1 <- total
}

func SumarServicios(servis []Servicios, c2 chan float64) {
	total := 0.0
	for _, servi := range servis {
		if servi.MinsTrab <= 30 {
			total += servi.Precio * 30
		} else {
			total += servi.Precio * servi.MinsTrab
		}
	}
	c2 <- total
}

func SumarMantenimiento(mantes []Mantenimiento, c3 chan float64) {
	total := 0.0
	for _, manti := range mantes {
		total += manti.Precio
	}
	c3 <- total
}
