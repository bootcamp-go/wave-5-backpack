package main

import "fmt"

//Definición de estructura

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

// Métodos de cambio utilizando punteros

func (u *Usuario) cambiarNombre(nombre *string, apellido *string) {
	u.Nombre = *nombre
	u.Apellido = *apellido
}

func (u *Usuario) cambiarEdad(edad *int) {
	u.Edad = *edad
}

func (u *Usuario) cambiarCorreo(correo *string) {
	u.Correo = *correo
}

func (u *Usuario) cambiarContrasena(password *string) {
	u.Contrasena = *password
}

func main() {

	usuario := Usuario{
		Nombre:     "NombreInicial",
		Apellido:   "ApellidoInicial",
		Edad:       15,
		Correo:     "correoInicial",
		Contrasena: "passwordIncial",
	}

	var (
		nombre   string = "José"
		apellido string = "Berrios"
		edad     int    = 20
		correo   string = "prueba@example.com"
		password string = "contrasena1234"
	)

	usuario.cambiarNombre(&nombre, &apellido)
	usuario.cambiarEdad(&edad)
	usuario.cambiarCorreo(&correo)
	usuario.cambiarContrasena(&password)
	fmt.Println(usuario)

}
