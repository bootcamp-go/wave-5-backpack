package main

import "fmt"

type User struct {
	nombre   string
	apellido string
	edad     int
	email    string
	password string
}

func main() {

	usuario := User{nombre: "Roko", apellido: "Doggy", edad: 3, email: "roko@dogmail.com", password: "doglife"}

	fmt.Println(usuario) //Usuario OG
	cambiarNombre(&usuario)
	cambiarApellido(&usuario)
	cambiarEdad(&usuario)
	cambiarEmail(&usuario)
	cambiarPassword(&usuario)

	fmt.Println(usuario) //Usuario 2.0

}

func cambiarNombre(u *User) {
	//fmt.Println(*&u.nombre)
	*&u.nombre = "Rocky"
}
func cambiarApellido(u *User) {
	//fmt.Println(*&u.nombre)
	*&u.apellido = "Dogger"
}
func cambiarEdad(u *User) {
	//fmt.Println(*&u.nombre)
	*&u.edad = 2
}
func cambiarEmail(u *User) {
	//fmt.Println(*&u.nombre)
	*&u.email = "rocky@dogmail.com"
}
func cambiarPassword(u *User) {
	//fmt.Println(*&u.nombre)
	*&u.password = "crazyDog"
}
