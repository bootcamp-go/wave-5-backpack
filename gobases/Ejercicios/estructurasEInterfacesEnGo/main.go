package main

import (
	"fmt"
)

func main() {
	/*//1
	alumno := Alumno{
		Nombre:       "Francisco",
		Apellido:     "Monay",
		DNI:          "35123321",
		FechaIngreso: "21/06/2022",
	}

	alumno.Detalle() */

	/*//2
	matriz := Matriz{
		Alto:       5,
		Ancho:      5,
		Cuadratica: true,
	}

	matriz.setMatriz(1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 2, 8, 7, 6, 5, 7, 8, 1, 0, 2, 2, 3, 4, 5)
	matriz.printMatriz()
	matriz.valorMax(matriz.Valores)
	*/

	//3
	productoPequeño := newProducto("pequeño", "Celular", 100)
	productoMediano := newProducto("mediano", "Notebook", 100.00)
	productoGrande := newProducto("grande", "Heladera", 100)

	tienda := tienda{}

	tienda.agregar(productoPequeño.calcularCosto())
	tienda.agregar(productoMediano.calcularCosto())
	tienda.agregar(productoGrande.calcularCosto())

	fmt.Println("El precio total de los productos es: ", tienda.total())

}
