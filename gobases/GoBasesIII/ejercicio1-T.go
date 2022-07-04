package main

import "fmt"

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Password string
}

func (u *Usuario) cambiarNombre(name, lastname string) (string, string) {
	u.Nombre = name
	u.Apellido = lastname

	return u.Nombre, u.Apellido
}

func (u *Usuario) cambiarEdad(age int) int {
	u.Edad = age
	return u.Edad
}

func (u *Usuario) cambiarCorreo(mail string) string {
	u.Correo = mail
	return u.Correo
}

func (u *Usuario) cambiarContrasena(pass string) string {
	u.Password = pass
	return u.Password
}

func main() {
	u := Usuario{
		Nombre:   "Juan",
		Apellido: "Serna",
		Edad:     23,
		Correo:   "juan.serna",
		Password: "12345",
	}

	fmt.Println(u.cambiarNombre("David", "Valderrama"))
	fmt.Println(u.cambiarEdad(20))
	fmt.Println(u.cambiarCorreo("serna.juan"))
	fmt.Println(u.cambiarContrasena("qwerty"))
}
