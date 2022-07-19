package main

import (
	"fmt"
	"os"
)

//Defino la estrucutura del producto
type Producto struct {
	ID       int
	precio   float64
	cantidad int
}

func main() {
	//Instancio una variable de tipo Producto
	listaProductos := []Producto{
		{111223, 30012.00, 1},
		{444321, 1000000.00, 4},
		{434321, 50.50, 1},
	}
	//Envio datos para guardar productos en CSV
	saveData(listaProductos)
}

func saveData(listaProductos []Producto) {
	csvInfo := "ID,Precio,Cantidad\n"
	for _, producto := range listaProductos {
		csvInfo += fmt.Sprintf("%d,%f,%d\n", producto.ID, producto.precio, producto.cantidad)
	}

	//Escribir archivo CSV
	err := os.WriteFile("productos.csv", []byte(csvInfo), 0644)
	if err != nil {
		fmt.Println("Error al sobreescribir el archivo")
	}
}
