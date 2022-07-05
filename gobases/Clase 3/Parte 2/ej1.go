package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func CambiarNombre(u *Usuario, nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func CambiarEdad(u *Usuario, edad int) {
	u.Edad = edad
}

func CambiarCorreo(u *Usuario, correo string) {
	u.Correo = correo
}

func CambiarContrasena(u *Usuario, contrasena string) {
	u.Contrasena = contrasena
}

func main() {
	user := Usuario{
		Nombre:     "Pepe",
		Apellido:   "Sierra",
		Edad:       15,
		Correo:     "valid@email.com",
		Contrasena: "secure_pass",
	}
	fmt.Println(user)
	pUser := &user
	CambiarNombre(pUser, "Pepito", "Perez")
	CambiarCorreo(pUser, "nuevo@correo.com")
	CambiarEdad(pUser, 16)
	CambiarContrasena(&user, "better_pass")
	fmt.Println(user)
}
