/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Red Social
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A university needs to register students and generate a functionality
		to print the details of each student's data, as follows:
			First name: [Student's first name]
			Surname: [Student's surname]
			Last name: [Student's surname
			ID: [Student's ID]
			Date: [Student's entry date]
		The values in square brackets must be replaced by the data provided
		by the students.
		For this it is necessary to generate a structure Students with the
		variables Name, Surname, ID, Date and with a detail method

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
)

//	STRUCT : Usr
type Usr struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
}

//	FUNCTIONS
func cambiarNombre(user *Usr, newName string) Usr {
	user.Nombre = newName
	return *user
}

func cambiarApellido(user *Usr, newLastName string) Usr {
	user.Apellido = newLastName
	return *user
}

func cambiarEdad(user *Usr, newAge int) Usr {
	user.Edad = newAge
	return *user
}

func cambiarCorreo(user *Usr, newMail string) Usr {
	user.Correo = newMail
	return *user
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Red Social ||")
	user1 := Usr{
		Nombre:   "Pancho",
		Apellido: "Lopez",
		Edad:     42,
		Correo:   "pancho.lopez@socialconnect.com",
	}

	fmt.Println("> Datos del usuario #1: ")
	fmt.Println("\t", user1.Nombre)
	fmt.Println("\t", user1.Apellido)
	fmt.Println("\t", user1.Edad)
	fmt.Println("\t", user1.Correo)

	fmt.Println("> Alterando los datos del usuario #1: ")
	user1 = cambiarNombre(&user1, "Chancho")
	user1 = cambiarApellido(&user1, "Sancho")
	user1 = cambiarEdad(&user1, 32)
	user1 = cambiarCorreo(&user1, "chancho.sancho@socialconnect.com")
	fmt.Println("\t", user1.Nombre)
	fmt.Println("\t", user1.Apellido)
	fmt.Println("\t", user1.Edad)
	fmt.Println("\t", user1.Correo)
}
