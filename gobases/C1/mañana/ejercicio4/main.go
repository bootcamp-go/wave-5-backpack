package main

import "fmt"

/*Ejercicio 4 - Tipos de datos

Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go pero tuvo varias dudas mientras lo hacía.
A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
Verificar su código y realizar las correcciones necesarias.
Corregir las incorrectas:
	var apellido string = "Gomez"
	var edad int = "35"
	boolean := "false";
	var sueldo string = 45857.90
	var nombre string = "Julián"
*/

func main() {
	var apellido string = "Gomez"

	// Utilizar el tipo de dato acorde a la información a utilizar -
	// Forma incorrecta var edad int = "35"
	var edad int = 35

	boolean := false

	// Utilizar el tipo de dato acorde a la información a utilizar -
	// Forma incorrecta var sueldo string =  45857.90
	var sueldo float32 = 45857.90

	var nombre string = "Julián"

	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}
