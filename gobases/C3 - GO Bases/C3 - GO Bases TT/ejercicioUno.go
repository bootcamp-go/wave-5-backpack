package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

//Implementación de funciones
func (u *Usuario) cambiarNombre(n, a string) {
	u.Nombre = n
	u.Apellido = a
}

func (u *Usuario) cambiarEdad(e int) {
	u.Edad = e
}

func (u *Usuario) cambiarCorreo(c string) {
	u.Correo = c
}

func (u *Usuario) cambiarContraseña(co string) {
	u.Contraseña = co
}

func imprimir(user Usuario) {
	fmt.Println("Nombre:", user.Nombre, "\nApellido:", user.Apellido, "\nEdad:", user.Edad)
	fmt.Println("Correo:", user.Correo, "\nContraseña:", user.Contraseña)
}

func main() {
	//Usuario ejemplo
	user := Usuario{
		Nombre:     "Luz Carime",
		Apellido:   "Lucumí",
		Edad:       24,
		Correo:     "luz.lucumi@correo.com.co",
		Contraseña: "qweratsQ@dhbhgeuyblablabla...",
	}

	//Información general
	fmt.Println("----------\nDatos de usuario inicial\n----------")
	imprimir(user)

	//Cambiando datos con los métodos
	user.cambiarNombre("Luz", "Lucumí Hernández")
	user.cambiarEdad(26)
	user.cambiarCorreo("luz.lucumi@hotmail.com")
	user.cambiarContraseña("Luz123")

	//Información después del cambio
	fmt.Println("----------\nDatos modificados del usuario\n----------")
	imprimir(user)
}
