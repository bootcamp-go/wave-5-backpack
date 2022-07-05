package main

import "fmt"

type Usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contrasena string
}

func setNombre(u *Usuario, nombre, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}

func setEdad(u *Usuario, edad int) {
	u.edad = edad
}

func setCorreo(u *Usuario, correo string) {
	u.correo = correo
}

func setContrasena(u *Usuario, contrasena string) {
	u.contrasena = contrasena
}

func main() {
	var p1 = new(Usuario)
	var p2 = new(Usuario)

	setNombre(p1, "Camilo", "Hernandez")
	setEdad(p1, 22)
	setCorreo(p1, "test@gmail.com")
	setContrasena(p1, "123")

	setNombre(p2, "Felipe", "Ruiz")
	setEdad(p2, 30)
	setCorreo(p2, "test2@gmail.com")
	setContrasena(p2, "123456")

	fmt.Println(p1)
	fmt.Println(p2)
}
