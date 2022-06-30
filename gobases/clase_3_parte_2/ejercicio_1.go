package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) CambiarNombre(nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) CambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) CambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) CambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main() {
	usuario := Usuario{
		Nombre:     "Claudio",
		Apellido:   "Figueroa",
		Edad:       25,
		Correo:     "claudio.figueroa@mercadolibre.cl",
		Contraseña: "contra",
	}

	fmt.Println(usuario)

	usuario.CambiarNombre("Juan", "Almirante")
	usuario.CambiarCorreo("juan.almirante@mercadolibre.cl")
	usuario.CambiarEdad(24)
	usuario.CambiarContraseña("juan232")

	fmt.Println(usuario)
}
