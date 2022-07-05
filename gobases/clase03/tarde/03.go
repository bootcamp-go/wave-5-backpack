/*
Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

package main
import "fmt"

type Productos struct {
	Nombre string
	Precio float64
	Cantidad int
}

type Servicios struct {
	Nombre string
	Precio float64
	minutosTrabajados int
}

type Mantenimientos struct {
	Nombre string
	Precio float64
}

func SumarProductos(productos *[]Productos, c chan float64){
	var total float64 = 0
	i :=1 // solo me sirve para imprimir la iteración y entender cómo funciona
	for _, producto := range *productos {
		total += producto.Precio * float64(producto.Cantidad)

		fmt.Println("Sumando Productos", i)// solo me sirve para imprimir la iteración y entender cómo funciona
		i++// solo me sirve para imprimir la iteración y entender cómo funciona
	}
	fmt.Printf("Total Prouctos: $%v\n", total)
	c <- total
	close(c)
}

func SumarServicios(servicios *[]Servicios, c chan float64) {
	var total float64 = 0
	i :=1 // solo me sirve para imprimir la iteración y entender cómo funciona
	for _, servicio := range *servicios {
		if servicio.minutosTrabajados < 30 {
			total += servicio.Precio * 30.0
		}else {
			total += servicio.Precio * float64(servicio.minutosTrabajados)
		}
		fmt.Println("Sumando Servicios", i)// solo me sirve para imprimir la iteración y entender cómo funciona
		i++// solo me sirve para imprimir la iteración y entender cómo funciona
	}
	fmt.Printf("Total Servicios: $%v\n", total)
	c <- total
	close(c)
	
}

func SumarMantenimientos(mantenimientos *[]Mantenimientos, c chan float64) {
	var total float64 = 0
	i :=1 // solo me sirve para imprimir la iteración y entender cómo funciona
	for _, mantenimiento := range *mantenimientos {
		total += mantenimiento.Precio
		fmt.Println("Sumando Mantenimiento", i)// solo me sirve para imprimir la iteración y entender cómo funciona
		i++// solo me sirve para imprimir la iteración y entender cómo funciona
	}
	fmt.Printf("Total Mantenimientos: $%v\n", total)
	c <- total
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
		{Nombre: "Programación", Precio: 200, minutosTrabajados: 480},
		{Nombre: "Paqueteria", Precio: 10234, minutosTrabajados: 30},
		{Nombre: "Encomienda", Precio: 300, minutosTrabajados: 15},
		{Nombre: "Limpieza", Precio: 100, minutosTrabajados: 22},
	}

	mantenimientos := []Mantenimientos{
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
	go SumarMantenimientos(&mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("* Total final * $ ", t1+t2+t3)

}

