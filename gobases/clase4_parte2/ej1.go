package main

import "fmt"

func main() {

	usuario := Usuario{"Armando", "Barreda", 54, "armando.barreda@gmail.com", "soldado_de_guerra"}

	fmt.Println("Usuario antes de modificarlo:", usuario)

	cambiarNombre(&usuario, "Simur", "Skinner")
	cambiarEdad(&usuario, 50)
	cambiarCorreo(&usuario, "simur.skinner@gmail.com")
	cambiarContrasenia(&usuario, "director_de_la_escuela")

	fmt.Println("Usuario antes de modificarlo:", usuario)
}

type Usuario struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func cambiarNombre(usuario *Usuario, nuevoNombre string, nuevoApellido string) {
	usuario.Nombre = nuevoNombre
	usuario.Apellido = nuevoApellido
}

func cambiarEdad(usuario *Usuario, nuevaEdad int) {
	usuario.Edad = nuevaEdad
}

func cambiarCorreo(usuario *Usuario, nuevoCorreo string) {
	usuario.Correo = nuevoCorreo
}

func cambiarContrasenia(usuario *Usuario, nuevaContrasenia string) {
	usuario.Contrasenia = nuevaContrasenia
}
