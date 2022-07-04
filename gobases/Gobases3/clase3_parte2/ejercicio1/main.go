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
		Apellido: "Salazar",
		Edad:     24,
		Correo:   "juansebastian.salazar@mercadolibre.com.co",
		Password: "56789",
	}

	fmt.Println(u.cambiarNombre("Sebastian", "Salazar"))
	fmt.Println(u.cambiarEdad(25))
	fmt.Println(u.cambiarCorreo("juansebastiansalaza@gmail.com"))
	fmt.Println(u.cambiarContrasena("juanse123"))
}
