package main

import "fmt"

//Ejercicio 1 - Red Social

type User struct {
	Nombre   string
	Apellido string
	Edad     int
	Email    string
	Password string
}

func nameChange(u *User, name string) {
	u.Nombre = name
	fmt.Println("el nombre ha sido modificado con exito")
}

func edadChange(u *User, edad int) {
	u.Edad = edad
	fmt.Println("la edad ha sido modificado con exito")
}

func emailChange(u *User, email string) {
	u.Email = email
	fmt.Println("el email ha sido modificado con exito")
}

func passwordChange(u *User, password string) {
	u.Password = password
	fmt.Println("la Contrasena ha sido modificado con exito")
}

func main() {
	andres := User{Nombre: "Andres",
		Apellido: "Ospina",
		Edad:     19,
		Email:    "andres@gmail.com",
		Password: "anuelito123"}

	fmt.Print(andres.Nombre, "\n")
	nameChange(&andres, "Joaquin")
	fmt.Print(andres.Nombre, "\n")

	fmt.Print(andres.Edad, "\n")
	edadChange(&andres, 20)
	fmt.Print(andres.Edad, "\n")

	fmt.Print(andres.Email, "\n")
	emailChange(&andres, "joaquin1@gmail.com")
	fmt.Print(andres.Email, "\n")
}
