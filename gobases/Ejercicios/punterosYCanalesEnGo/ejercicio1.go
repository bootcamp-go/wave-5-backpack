package main

import "fmt"

type usuario struct {
	Nombre    string
	Apellido  string
	Edad      int
	Correo    string
	Password  string
	Productos []producto
}

func (u usuario) detalleU() {
	fmt.Printf("Nombre: %s\nApellido: %s\nEdad: %d\nCorreo: %s\nPassword: %s\n", u.Nombre, u.Apellido, u.Edad, u.Correo, u.Password)
}

func newUsuario(nombre string, apellido string, edad int, correo string, password string) usuario {
	return usuario{
		Nombre:   nombre,
		Apellido: apellido,
		Edad:     edad,
		Correo:   correo,
		Password: password,
	}
}

func (u *usuario) cambiarNombre(nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *usuario) cambiarPassword(password string) {
	u.Password = password
}
