package main

import "fmt"

func main() {
	/*

		var apellido string = "Gomez"
		var edad int = "35"
		boolean := "false";
		var sueldo string = 45857.90
		var nombre string = "Julián"

	*/

	var apellido string = "Gomez"
	// var edad int = "35" : error por usar comillas en un int
	// boolean := "false" : error por usar comillas en false
	//var sueldo string = 45857.90 : error por usar un float en un string
	var nombre string = "Julian"

	fmt.Println(apellido, nombre)

}
