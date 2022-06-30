package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func cambiarNombre(u *usuario, nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func cambiarEdad(u *usuario, edad int) {
	u.Edad = edad
}

func cambiarCorreo(u *usuario, correo string) {
	u.Correo = correo
}

func cambiarContrasena(u *usuario, contrasena string) {
	u.Contrasena = contrasena
}

func main() {
	u1 := usuario{"Jessica", "Escobar", 22, "jessica@hotmail.com", "1234"}
	fmt.Println("Incial: ", u1)

	cambiarNombre(&u1, "Paola", "Perez")
	cambiarEdad(&u1, 23)
	cambiarCorreo(&u1, "paola@hotmail.com")
	cambiarContrasena(&u1, "5678")
	fmt.Println("Cambios: ", u1)
}
