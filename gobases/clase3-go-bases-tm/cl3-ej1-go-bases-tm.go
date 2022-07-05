/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Guardar archivo
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A company that sells cleaning products needs:
			1.Implement a functionality to store a text file, with the information
			  of products purchased, separated by semicolons (csv).
			2. It must have the product id, price and quantity.
			3. These values can be hardcoded or hard written in a variable.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
	"os"
)

//	STRUCT : Producto
type Producto struct {
	ID       int
	PRECIO   float64
	CANTIDAD int
}

//	FUNCTION : guardarArchivo
func guardarArchivo(lista ...Producto) error {
	var s string
	archivo := "productos.csv"
	if archivoExiste(archivo) { // En caso de existir el archivo CSV
		f, err := os.OpenFile(archivo, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		for _, prod := range lista {
			s += fmt.Sprintf("%d,%.2f,%d\n", prod.ID, prod.PRECIO, prod.CANTIDAD)
		}
		fmt.Fprintf(f, "\n%s\n", s)
		defer f.Close()
		return err
	} else { // En caso de NO existir el archivo CSV
		s = "ID,PRECIO,CANTIDAD\n"
		for _, prod := range lista {
			s += fmt.Sprintf("%d,%.2f,%d", prod.ID, prod.PRECIO, prod.CANTIDAD)
		}
		data := []byte(s)
		err := os.WriteFile("./productos.csv", data, 0644) // Escribe y Guarda txt
		return err
	}
}

//	FUNCTION : archivoExiste
func archivoExiste(route string) bool {
	if _, err := os.Stat(route); os.IsNotExist(err) {
		return false
	}
	return true
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Guardar en un Archivo ||")

	// Crea archivo CSV & guarda el producto
	guardarArchivo(Producto{ID: 1, PRECIO: 100, CANTIDAD: 2})

	// Actualiza el archivo CSV & guarda los productos
	guardarArchivo(
		Producto{ID: 2, PRECIO: 200, CANTIDAD: 50},
		Producto{ID: 3, PRECIO: 150, CANTIDAD: 24},
	)

	fmt.Println("Archivo guardado como 'productos.csv'")
}
