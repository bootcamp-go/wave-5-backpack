/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Datos de clientes
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Repeat the previous process, but now implementing " fmt.Errorf() ",
		so that the error message receives by parameter the value of " salary "
		indicating that it does not reach the taxable minimum (the message displayed
		by console should read: " error: the taxable minimum is 150,000 and
		the salary entered is: [salary] ", where [ salary ] is the value of
		type int passed by parameter.
	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
	"os"
)

//	FUNCTION : leerArchivo
func leerArchivo(nameFile string) *os.File {
	fileData, err := os.Open(nameFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ejecución Finalizada")
	return fileData
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Datos de Clientes ||")

	data := leerArchivo("customers.txt")
	fmt.Println(data)
}
