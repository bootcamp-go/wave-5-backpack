package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (usuario *Usuario) ActualizarNombre(nombre *string) {
	usuario.Nombre = *nombre
}
func (usuario *Usuario) ActualizarApellido(apellido *string) {
	usuario.Apellido = *apellido
}
func (usuario *Usuario) ActualizarEdad(edad *int) {
	usuario.Edad = *edad
}
func (usuario *Usuario) ActualizarCorreo(correo *string) {
	usuario.Correo = *correo
}
func (usuario *Usuario) ActualizarContrasena(contrasena *string) {
	usuario.Contrasena = *contrasena
}

func main() {
	u := &Usuario{
		Nombre:     "Luz",
		Apellido:   "Jimenez",
		Edad:       24,
		Correo:     "ligth@gamil.com",
		Contrasena: "1234567",
	}

	var (
		nombre     string = "Elena"
		apellido   string = "Urrego"
		correo     string = "elena@gmail.com"
		edad       int    = 25
		contrasena string = "12345"
	)

	u.ActualizarNombre(&nombre)
	u.ActualizarApellido(&apellido)
	u.ActualizarCorreo(&correo)
	u.ActualizarEdad(&edad)
	u.ActualizarContrasena(&contrasena)

	fmt.Printf("Nombre: %s\nEdad: %d\nCorreo: %s\nConstrasena: %s\n", u.Nombre, u.Edad, u.Correo, u.Contrasena)

}
