// Ejercicio 3 - Calcular Precio
// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

// Ejercicio 4 - Ordenamiento
// Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
// Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
// un arreglo de números enteros con 100 valores
// un arreglo de números enteros con 1000 valores
// un arreglo de números enteros con 10000 valores

// Para instanciar las variables utilizar rand
// /* package main
 
// import (
//    "math/rand"
// )
 
 
// func main() {
//    variable1 := rand.Perm(100)
//    variable2 := rand.Perm(1000)
//    variable3 := rand.Perm(10000)
// }
//  */
// Se debe realizar el ordenamiento de cada una por:
// Ordenamiento por inserción
// Ordenamiento por burbuja
// Ordenamiento por selección

// Una go routine por cada ejecución de ordenamiento
// Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
// Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo

package main

import "fmt"
type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad uint64
}


func SumarProducto(productos *[]Producto, c chan float64) {
	var total float64
	for _, value := range *productos {
		total += value.Precio * float64(value.Cantidad)
	}

	fmt.Println("\nTotal productos $", total)
	c <- total
	close(c)
}



func main() {
	prod := []Producto{
		{Nombre: "p1", Precio: 123, Cantidad: 5},
		{Nombre: "p2", Precio: 321, Cantidad: 8},
		{Nombre: "p3", Precio: 234, Cantidad: 33},
		{Nombre: "p4", Precio: 432, Cantidad: 7},
	}
	canal1 := make(chan float64)

	go SumarProducto(&prod, canal1)
	fmt.Println("\naca van los productos ", prod)
}
