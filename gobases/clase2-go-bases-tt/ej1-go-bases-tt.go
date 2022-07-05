/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Registro de estudiantes
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A university needs to register students and generate a functionality to
		print the details of each student's data, as follows:
			First name: [Student's first name]
			Surname: [Student's surname] Last name: [Student's surname
			ID: [Student's ID] Date: [Student's entry date]
			Date: [Student's entry date]
		The values in square brackets must be replaced by the data provided
		by the students.
		For this it is necessary to generate a structure Students with
		the variables Name, Last name, ID, Date and that has a detail method
	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	STRUCT : Alumnos
type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

//	FUNCTION : al.detalle()
func (al Alumnos) detalle() {
	fmt.Printf("\n Nombre:\t%s\n Apellido:\t%s\n DNI:\t\t%d\n Fecha:\t\t%s\n", al.Nombre, al.Apellido, al.DNI, al.Fecha)
}

//	MAIN PROGRAM
func main() {
	al1 := Alumnos{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		DNI:      834892374983,
		Fecha:    "06-07-1997",
	}
	fmt.Println("\n\t|| Registro de Estudiantes ||")
	al1.detalle()
}
